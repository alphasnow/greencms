package main

import (
	"context"
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"server/api/swagdocs"
	router2 "server/internal/admin/router"
	"server/internal/core/model/dao"
	"server/internal/core/model/utils"
	"server/internal/web/router"
	"server/pkg/g"
	"time"
)

var envMode string

func init() {
	// 注意这里我们使用了字符串类型的标志
	flag.StringVar(&envMode, "env", "", "set the environment (e.g., dev, prod)")
	flag.Parse()
}

// main
// @title					API接口文档
// @description				遵循Restful API接口规范
// @version					1.0.0
// @contact.name AlphaSnow
// @contact.email wind91@foxmail.com
// @basePath					/
// @tag.name Web
// @tag.description 前端数据服务接口
func main() {
	db := g.DB()
	dao.SetDefault(db)
	if err := utils.AutoMigrate(db); err != nil {
		log.Fatalf("db auto migrate error: %s", err)
	}
	if err := utils.AutoInitialData(db); err != nil {
		log.Fatalf("db auto initial data error: %s", err)
	}

	cfg := g.Config()
	if cfg.GetBool("app.debug") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
		// 错误信息打印在屏幕
		// gin.DefaultErrorWriter = io.Discard
	}

	if envMode != "" {
		cfg.Set("app.env", envMode)
	} else {
		envMode = cfg.GetString("app.env")
	}

	servers := make([]*http.Server, 0)
	errNotify := make(chan error, 1)

	if envMode == "prod" && cfg.GetBool("server.web.status") {
		go func() {
			ser := buildWebServer(cfg)
			servers = append(servers, ser)
			errNotify <- ser.ListenAndServe()
		}()
		log.Printf("启动前台服务 %s\n", cfg.GetString("server.web.url"))
	}
	if envMode == "prod" && cfg.GetBool("server.admin.status") {
		go func() {
			ser := buildAdminServer(cfg)
			servers = append(servers, ser)
			errNotify <- ser.ListenAndServe()
		}()
		log.Printf("启动后台服务 %s\n", cfg.GetString("server.admin.url"))
	}
	if cfg.GetBool("server.api.status") {
		go func() {
			svc := buildApiServer(cfg)
			servers = append(servers, svc)
			errNotify <- svc.ListenAndServe()
		}()
		log.Printf("启动后端服务 %s\n", cfg.GetString("server.api.url"))
	}

	// 若判断当前是window系统自动打开前端网页
	if runtime.GOOS == "windows" {
		if g.Config().GetBool("server.api.open") {
			_ = browser.OpenURL(cfg.GetString("server.api.url"))
		}
		if envMode == "prod" && g.Config().GetBool("server.admin.open") {
			_ = browser.OpenURL(cfg.GetString("server.admin.url"))
		}
		if envMode == "prod" && g.Config().GetBool("server.web.open") {
			_ = browser.OpenURL(cfg.GetString("server.web.url"))
		}
	}

	quickNotify := make(chan os.Signal, 1)
	signal.Notify(quickNotify, os.Interrupt, os.Kill)
	select {
	case err := <-errNotify:
		log.Printf("服务发生错误 %s\n", err)
	case sg := <-quickNotify:
		log.Printf("服务主动退出 %s\n", sg)
	}

	log.Println("服务等待关闭 ...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	finish := make(chan struct{}, 1)
	go func() {
		for _, s := range servers {
			if err := s.Shutdown(ctx); err != nil {
				log.Printf("服务关闭发生错误 %s\n", err)
			}
		}
		finish <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		log.Fatalln("服务关闭超时")
	case <-finish:
		log.Println("服务关闭完成")
	}
}

func buildWebServer(cfg *viper.Viper) *http.Server {
	r := gin.Default()
	// 开启后JS被压缩 173kb 减少为54.2kb
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Static("/", g.Path("website/web"))
	s := &http.Server{
		Addr:    cfg.GetString("server.web.addr"),
		Handler: r,
	}
	return s
}
func buildAdminServer(cfg *viper.Viper) *http.Server {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Static("/", g.Path("website/admin"))
	s := &http.Server{
		Addr:    cfg.GetString("server.admin.addr"),
		Handler: r,
	}
	return s
}
func buildApiServer(cfg *viper.Viper) *http.Server {
	r := gin.Default()

	// cors只能用在全局,因为需要处理全局的option请求
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		MaxAge:           2 * time.Hour,
	}))

	// requestid
	r.Use(requestid.New())

	// logger
	// logger access
	//r.Use(ginzap.GinzapWithConfig(g.Log(), &ginzap.Config{
	//	UTC:        true,
	//	TimeFormat: time.RFC3339,
	//	SkipPathRegexps: []*regexp.Regexp{
	//		regexp.MustCompile(`^/docs`),
	//		regexp.MustCompile(`^/debug`),
	//		regexp.MustCompile(`^/upload`),
	//		regexp.MustCompile(`^/api/admin`),
	//	},
	//}))
	// logger recovery
	r.Use(ginzap.RecoveryWithZap(g.Log("server"), true))

	// timeout
	// 若响应时间超过1秒,强制超时
	// 后期改为桶队列
	//r.Use(timeout.New(
	//	timeout.WithTimeout(1000*time.Microsecond),
	//))

	if g.Config().GetBool("server.api.docs") {
		//docs
		//godocs.SwaggerInfo.Title = "API接口说明文档"
		//godocs.SwaggerInfo.Description = "可使用Ctrl+F按路径搜索接口"
		//godocs.SwaggerInfo.Version = "1.0"
		//godocs.SwaggerInfo.BasePath = "/"
		urlAddr, _ := url.Parse(g.Config().GetString("server.api.url"))
		swagdocs.SwaggerInfo.Host = urlAddr.Host
		swagdocs.SwaggerInfo.Schemes = []string{"http"}
		// docs
		r.GET("/docs", func(c *gin.Context) {
			c.Redirect(301, "/docs/index.html")
		})
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// API首页重定向到docs
		r.GET("/", func(c *gin.Context) {
			c.Redirect(301, "/docs/index.html")
		})
	}

	if g.Config().GetBool("app.debug") {
		// /debug/pprof
		pprof.Register(r)
	}

	// 更复杂可考虑使用 https://github.com/gin-contrib/static
	r.Static("/upload", g.Path("storage/upload"))

	router.Register(r)
	router2.Register(r)

	s := &http.Server{
		Addr:           cfg.GetString("server.api.addr"),
		Handler:        r,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		IdleTimeout:    1 * time.Second, // 空闲连接1秒后关闭
		MaxHeaderBytes: 1 << 20,         // 限制请求头大小1M以内
	}
	return s
}
