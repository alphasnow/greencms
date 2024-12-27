package g

import (
	"github.com/spf13/viper"
	"path"
	"server/pkg/xconfig"
)

func Config() (cfg *viper.Viper) {
	var err error
	if globalContainer.Has("config") {
		cfg = globalContainer.Get("config").(*viper.Viper)
	} else {
		// cfg, err = xconfig.ReadConfig(Path("config.yaml"))
		cfg, err = xconfig.ReadConfigAndEnv(Path("config.yaml"), Path(".env"))
		if err != nil {
			panic(err)
		}

		toAbsPath(cfg)
		globalContainer.Set("config", cfg)
	}
	return cfg
}

func toAbsPath(cfg *viper.Viper) {
	needAbs := []string{
		"server.api.path",
		"server.admin.path",
		"server.web.path",
		"database.connections.sqlite.filename",
		"logger.app.filename",
		"logger.server.filename",
	}
	for _, v := range needAbs {
		p := cfg.GetString(v)
		if path.IsAbs(p) {
			continue
		}

		cfg.Set(v, Path(p))
	}

}
