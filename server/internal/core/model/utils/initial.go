package utils

import (
	"context"
	"encoding/json"
	// "github.com/samber/lo"
	"gorm.io/gorm"
	"io"
	"io/fs"
	"log"
	"math/rand/v2"
	"os"
	"path/filepath"
	"server/internal/core/model/accessor"
	"server/internal/core/model/dao"
	"server/internal/core/model/entity"
	"server/pkg/g"
	"server/pkg/utils"
	"strings"
	"time"
)

func AutoInitialData(db *gorm.DB) (err error) {
	locker := g.Path("storage/data/initial.lock")
	if _, err = os.Stat(locker); err == nil {
		return nil
	}
	defer func() {
		if err != nil {
			// log.Println("数据初始化失败")
			return
		}

		_ = os.WriteFile(locker, []byte(time.Now().Format(time.DateTime)), 0666)
	}()
	log.Println("数据初始化")

	c := context.Background()

	// default data
	passStr := getDefaultAdminPassword()
	log.Printf("后台登录密码 %s \n", passStr)
	pass, _ := utils.PasswordHash(passStr)

	if err = createDefaultAdminUser(c, pass); err != nil {
		return err
	}
	if err = createDefaultWebMeta(c); err != nil {
		return err
	}

	// demo data
	//if g.Config().GetString("app.env") == "dev" {
	if err = createDemoAdminUser(c, pass); err != nil {
		return err
	}
	if err = createDemoWebBanner(c); err != nil {
		return err
	}
	if err = createDemoArticleCategory(c); err != nil {
		return err
	}

	if err = createDemoArticleTag(c); err != nil {
		return err
	}

	if err = createDemoArticle(c); err != nil {
		return err
	}
	//}

	// log.Println("数据初始化完成")
	return nil
}

func getDefaultAdminPassword() string {
	var passStr string
	//if g.Config().GetString("app.env") == "dev" {
		passStr = "admin"
	// } else {
	// 	// 随机8位密码
	// 	passStr = lo.RandomString(8, lo.AlphanumericCharset)
	// }
	return passStr
}

func createDefaultAdminUser(c context.Context, pass string) error {
	// 若没有admin则填充默认
	_, err := dao.AdminUser.WithContext(c).Where(dao.AdminUser.ID.Eq(1)).Take()
	if err == nil {
		return nil
	}

	f, _ := os.OpenFile(g.Path("storage/app/example/admin_users.json"), os.O_RDONLY, 0)
	defer f.Close()
	fb, _ := io.ReadAll(f)
	var users []*entity.AdminUser
	_ = json.Unmarshal(fb, &users)

	for _, v := range users {
		v.Password = pass
	}

	if err = dao.AdminUser.WithContext(c).Create(users[0]); err != nil {
		return err
	}
	return nil
}

func createDemoAdminUser(c context.Context, pass string) error {
	fj := "storage/app/example/admin_users.json"
	f, _ := os.OpenFile(g.Path(fj), os.O_RDONLY, 0)
	defer f.Close()
	fb, _ := io.ReadAll(f)

	var users []*entity.AdminUser
	_ = json.Unmarshal(fb, &users)

	for _, v := range users {
		v.Password = pass
	}

	users = users[1:]
	if err := dao.AdminUser.WithContext(c).Create(users...); err != nil {
		return err
	}

	return nil
}

func createDefaultWebMeta(c context.Context) error {
	fj := "storage/app/example/web_metas.json"
	f, _ := os.OpenFile(g.Path(fj), os.O_RDONLY, 0)
	defer f.Close()
	fb, _ := io.ReadAll(f)

	var data []*entity.WebMeta
	_ = json.Unmarshal(fb, &data)

	err := dao.WebMeta.WithContext(c).Create(data...)
	return err
}

func createDemoWebBanner(c context.Context) error {
	fj := "storage/app/example/web_banners.json"
	f, _ := os.OpenFile(g.Path(fj), os.O_RDONLY, 0)
	defer f.Close()
	fb, _ := io.ReadAll(f)

	var data []*entity.WebBanner
	_ = json.Unmarshal(fb, &data)

	err := dao.WebBanner.WithContext(c).Create(data...)
	return err
}

func createDemoArticleCategory(c context.Context) error {

	fj := "storage/app/example/article_categories.json"
	f, _ := os.OpenFile(g.Path(fj), os.O_RDONLY, 0)
	defer f.Close()
	fb, _ := io.ReadAll(f)

	var data []*entity.ArticleCategory
	_ = json.Unmarshal(fb, &data)

	err := dao.ArticleCategory.WithContext(c).Create(data...)
	return err
}

func createDemoArticleTag(c context.Context) error {

	fj := "storage/app/example/article_tags.json"
	f, _ := os.OpenFile(g.Path(fj), os.O_RDONLY, 0)
	defer f.Close()
	fb, _ := io.ReadAll(f)

	var data []*entity.ArticleTag
	_ = json.Unmarshal(fb, &data)

	err := dao.ArticleTag.WithContext(c).Create(data...)
	return err
}
func createDemoArticle(c context.Context) error {
	fp := "storage/app/example/articles"
	fjs, _ := getJSONFiles(g.Path(fp))
	for _, fj := range fjs {
		data := entity.Article{}
		dataContent := entity.ArticleContent{}
		{
			f, _ := os.OpenFile(fj, os.O_RDONLY, 0)
			defer f.Close()
			fb, _ := io.ReadAll(f)
			_ = json.Unmarshal(fb, &data)
		}
		{
			fjc := strings.Replace(fj, ".json", ".html", 1)
			f, _ := os.OpenFile(fjc, os.O_RDONLY, 0)
			defer f.Close()
			fb, _ := io.ReadAll(f)
			dataContent.Content = string(fb)
		}
		data.ArticleContent = &dataContent
		data.ArticleStatistic = &entity.ArticleStatistic{
			Views:      rand.UintN(500) + 1,
			Favourites: rand.UintN(50) + 1,
		}
		data.ArticleTags = []*entity.ArticleTag{
			{ID: rand.UintN(9) + 1},
			{ID: rand.UintN(9) + 1},
			{ID: rand.UintN(9) + 1},
		}
		_ = dao.Article.WithContext(c).Create(&data)
	}

	return nil
}

func getJSONFiles(folderPath string) ([]string, error) {
	var jsonFiles []string
	err := filepath.Walk(folderPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			jsonFiles = append(jsonFiles, path)
		}
		return nil
	})
	return jsonFiles, err
}

func createDemoArticles(c context.Context) error {
	content1 := `
<h1 class="pgc-h-arrow-right">导语</h1>
<p>近年来，我国的风电发电装机容量持续增长，新建风电项目的投建数量位居世界前列，在世界上具备很强的竞争力。</p>
<p>风电具有资源广泛、环保美观、安全稳定等特点，因此被广泛推广应用，有利于实现绿色低碳的生产模式。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-axegupay5k/92b7440eb942456eac2b0c7f69fad15b~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=PFBMeq%2BuPhaPSvIIC1wa%2FjB%2Btwg%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p><strong>但是我国在风电的发展领域中，还存在一些问题，需要通过技术创新等方式进行突破，帮助我国的风电产业更上一层楼。</strong></p>
<p>建设风力发电机，需要花多少钱?一天能够发多少电?以及多久能够回本?</p>
<h1 class="pgc-h-arrow-right">价格。</h1>
<p>风力发电机属于可再生能源的一种，其发电的原理是通过风力推动风轮，然后风轮与发电机进行链接，将磁力转换为电能。</p>
<p><strong>而风轮是风力发电机的关键部件，不同风力发电机的风轮也有所不同，根据风力发电机设计排水适用风速不同，需要的风轮尺寸也不一样。</strong></p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/7d54cf6c8fa542bcbfed3878ad2ea6ce~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=KYPR%2FxWzfq2LpJzUE4k8Ucy9xls%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>风轮的价格在3~6万之间，小尺寸的风轮价格为3万人民币，而中等规模的风轮价格在5万人民币，大尺寸的风轮价格在6万人民币。</p>
<p>风轮的价格是与其叶片的数目和长度有关，叶片的数目越多，综合性能也越好，同等条件下，叶片的长度越长，风轮所产生的电量也越大。</p>
<p>所以，一般风力发电机的叶片的数目在3~6片之间，叶片的长度一般在6~50米之间，而叶片的价格在2.8万~5.6万之间</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/a0e0c1475ed441e6a2cd5f08c22060ac~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=V4SzvI5q5NvH3Rk%2BUZoM0JHxOKw%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>发电机是将风轮的功率输出，转化为电能输出的设备，发电机的价格一般在2万~15万人民币之间，价格上升的原因主要是输出功率较大。</p>
<p>另外一方面，发电机的转速也在500~4000转/分钟之间，转向的轴也有很多种，例如水平轴，垂直轴。</p>
<p>发电机的价格主要取决于需要的功率大小和转速，不同功率大小的发电机价格也有所不同。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/1318d8da9c464bbbb866ca689f720bdc~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=JxK9eHm1BNMDSAwpl5ISzB0UPT0%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p><strong>塔架是通过风力发电机的高度，降低空气密度，提高风能利用率，因此，塔架高度一般在20米~100米之间，价格在1.8万~7.5万人民币之间。</strong></p>
<p>整个风力发电系统的价格大概在10万~80万人民币之间，但是，还需要考虑挖掘基础的价格，整个风力发电系统的价格大概在100万~500万人民币之间。</p>
<h1 class="pgc-h-arrow-right">风电发电技术。</h1>
<p>我国风能资源十分丰富，分布非常广泛，大部分地区都具备搭建风电场的条件，且我国的风能资源不仅陆地上具备，还十分丰富，我国海域也具备丰富的风能资源。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/cee19990f8d64f9ab0cfa5440216d19a~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=%2FlrU%2F03Jpn2jJIBXV5giDHDHGQc%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>我国在全球的风电技术领域中，处于领先地位，拥有丰富的风力发电机的制造经验，因此，我国在风电技术领域占有有利的一席之地。</p>
<p>在探索风能资源的过程中，我国根据当地的风能资源状况，以及环境情况，分别建立起了陆地风电场和海上风电场，形成了风力发电产业链。</p>
<p><strong>在风电技术方面，我国的风电技术与国外的风电技术相比，起步较晚，技术的水平可能相对来说较为低级一些。</strong></p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/1fc53fdfe0f04fecbef134e33fce90a4~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=3I%2BDdSHNxEHnAHZCyGOg%2B3hOQf8%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>在风电技术的发展过程中，我国进行了技术创新，以及吸收国外的一些先进技术，利用先进的技术对其进行改进，在技术发展速度上，也比较快，取得了很大的进步。</p>
<p><strong>但是在风电技术方面，我国对其进行了研究探索，建立风力发电机的样机，以及进行大规模的风电场建设，但是对其核心技术并没有过多的研究。</strong></p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/8a9cfdb8b6ea44c19d2d7534b8dd7be3~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=Jb7%2B3ff8UbjdUNpD4vIffASBn3Q%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>在我国建设风电场过程中，风力发电系统，是从国外引进来的，对于核心技术并没有进行创新，大部分发电系统的核心技术，仍然掌握在国外的企业手中。</p>
<p>风电发电机是风电系统中的核心部件，由于这部分的核心技术掌握在国外企业手中，因此，风电发电机的价格一直处于上涨的趋势中，成为制约我国风电产业发展的一大问题。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p26-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/0c001a906dc946c1bb841bb915ce4f71~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=Lkc2WxjZqjNT27GEG5pMls%2BoYhY%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>我国在风电技术方面的研究探索是十分深入的，对于风力发电机的研究也是十分深入的，我国研究人员进行了很多探讨。</p>
<p>风力发电机在生产制造的时候，需要进行铸造工艺，但是这一工艺比较复杂，生产制造工艺比较成熟，风力发电机的制造出来的时候，往往是一个独特存在。</p>
<p>我国在材料技术方面的发展是十分迅猛的，材料方面有很多成熟的方案，对于风力发电机在制造过程的工艺，我国研究人员进行了很多的尝试。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/c7efaef698de43aeb2b55d7c156e3642~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=CGVUp9MJfIpVLH1nuvUS7q5P2u8%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>在风力发电机上，我国研究人员进行了反复催熟的尝试，将其进行模块化设计，拼接组装，从而提高生产效率，降低制造成本。</p>
<p>风力发电机在生产制造过程中，如果能够进行模块化设计，能够进行一体化的组装，就能大大的降低生产周期，降低人工成本，进一步降低制造成本，提高生产效率。</p>
<p>风力发电机在制造过程中，如果进行模块化的拼接，并且采用多工位同时进行生产，就能够大幅度的提高生产效率。</p>
<h1 class="pgc-h-arrow-right">风电的发展。</h1>
<p>风电在我国的发展史上，起步较晚，但是在发展过程中，发展速度较快，新建风电项目的投资也在不断的增大。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p9-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/2f8ffe25864c4add88aff5126f125376~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=NkrGeMVn6MCwRzOpq5M318J75xo%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>我国风电在发展过程中，发展速度较快，已经在风电领域中，占据了有利的一席之地，成为世界上风电装机容量最多的国家。</p>
<p>但是，在风电领域领先的位置，我国还是需要在核心技术方面进行研究创新，对于风力发电机的核心技术，我们应该进行研究创新，争取能够进一步提高国产风电的安全性和稳定性。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/7219a8e2607c4d5c9c6c8f4d0f988738~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=wzuDolVZL2wiNNTjdzDFBJRpnug%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p><strong>风电在我国的发展过程中，对于环境的保护以及资源的利用，都具有重要的意义，在未来，我国风电领域还将迎来新的突破。</strong></p>
<h1 class="pgc-h-arrow-right">风力发电的收益。</h1>
<p>风电发电机在发电的同时，也会产生一定的噪音，对于噪音会带来一定的污染，但是在风电发电机的发展过程，对于噪音的处理也比较成熟。</p>
<p><strong>风电在我国的发展过程中，会对当地的经济产生一定的影响，风电项目一旦投资建设，还会带动附近的一些配套设施的建设。</strong></p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/70560505870d4d67be7744d94576ea78~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=EZV5Sk8sq6YfcH%2FEY1tFQ5UoDlU%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p><strong>在风电项目建设过程中，对于很多的当地村民，都有很多的收益，风电项目建设后，当地一些农民租地获得的一些收益，这种租地金对于一些农民来说，也是比较可观的。</strong></p>
<p>风电项目建设起来后，还需要一些工人进行清洗，对于一些当地的农民进行了招聘，从而解决了农民的一些就业问题。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/9ae6042864694e339162af7b6f6f14cf~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=ZgjMglaDrIlw%2B%2Fo9y27k22IVoAI%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>风电项目建设完成后，还会对当地的环境产生一定的影响，风电产生的电能清洁无污染，对于当地的环境产生的影响不大，对于我国碳排放的减少，起到了很好的作用。</p>
<h1 class="pgc-h-arrow-right">风力发电的回本时间。</h1>
<p>风力发电机的回本时间，主要取决于当地的风力资源状况，以及发电机的投资金额，以及每天的发电量。</p>
<p>风力发电在回本的过程中，需要较长的时间，产生的电价和成本价相比，产生的电价是比较低的，这也是风力发电回本时间比较长的原因。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/0396e6037bd34299a844d0fe30ff85f6~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=Ork4K4u8jYTzxOx4XGhMJY1vgmQ%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<p>我国在风电领域的发展势头十分强劲，在风电领域的发展过程中，风电产业链也会不断的完善，我国在风力发电机的制造过程中，逐渐掌握核心技朧，从而减少对国外进口的依赖。</p>
<p>风力发电机在我国的发展过程中，风电产业链也会逐渐形成，风力发电的核心技术逐渐掌握，风力发电机的核心技术掌握在我国的手中，风力发电机的价格也会得到降低。</p>
<div class="pgc-img"><img class="syl-page-img" src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/8d14df5e23a14901b72b336ecc6d2d70~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1722930424&amp;x-signature=DK69O770fKsb8fZV3CdhI8ZoY2k%3D">
<p class="pgc-img-caption">&nbsp;</p>
</div>
<h1 class="pgc-h-arrow-right">结语</h1>
<p><strong>风力发电机在发电的过程中，不仅能够保护我们的环境，还能够进行资源的开发利用，对于我国风电产业也具备重要的意义，我国在风电领域中也有着很大的发展前景。</strong></p>
`
	content2 := `<article class="syl-article-base tt-article-content syl-page-article syl-device-pc"><p data-track="1">微信只清空聊天记录等于没删，教你用正确清理方法能清理好几个G</p><p data-track="2">身为一个在互联网上冲浪多年的老鸟，肯定有很多人跟我一样，手机的内存常年告急。而微信，这个社交巨头，往往是内存占用的“大户”。大家都知道，聊天记录是微信内存消耗的大头。为了腾出点空间，我们经常会选择清空聊天记录。但是，如果只是简单地点一下“清空”按钮，你会发现，微信占用的内存并没有明显减少。</p><div class="pgc-img"><img src="https://p3-sign.toutiaoimg.com/tos-cn-i-axegupay5k/6fc0518e1b4540e8939c60ea6944f9ec~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=ii4upkU070I6jXDcFlLLzRHHJwE%3D" img_width="500" img_height="889" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/52a8db2590014c69bb3aa56b31e2fed9" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div><p data-track="3">这是因为微信有一个“小聪明”：它会把一些重要的聊天记录保存在其他地方，而不仅仅是聊天记录列表里。所以，仅仅清空聊天记录列表，并不能真正释放大量内存。</p><p data-track="4">那么，怎么才能真正清理微信，释放大量内存呢？别着急，我来教你几招。</p><div class="pgc-img"><img src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/9ef189c0bc834c608a259d8c12ead187~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=Q2bP3QaVG3I7q4SRKCAQN%2BYx9kA%3D" img_width="500" img_height="889" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/9ef189c0bc834c608a259d8c12ead187" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div><p data-track="5">一、深度清理微信缓存</p><p data-track="6">1. 进入微信存储空间使用情况</p><p data-track="7">首先，打开微信，点击右下角的“我”，然后选择“设置”。在设置页面，找到“存储使用情况”或者“数据和管理使用情况”（不同版本的微信可能叫法略有不同）。</p><div class="pgc-img"><img src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/fcd9caa4db3f450ca774b70c6ae86319~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=lvHI1opdMQlOvOMzwDTAlAzTLQo%3D" img_width="1080" img_height="1920" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/fcd9caa4db3f450ca774b70c6ae86319" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div><p data-track="8">2. 查看各个功能的存储使用情况</p><p data-track="9">在“存储使用情况”页面，你会看到微信各个功能的存储使用情况，包括“存储使用”、“管理…”等选项。点击“存储使用”，你会看到各个功能的详细存储使用情况。</p><p data-track="10">3. 手动清理缓存</p><p data-track="11">在众多功能中，点击“存储使用”较多的那个（通常是“微信”或者“XXX缓存”），然后点击“存储使用”下面的“管理…”。在这个页面，你会看到很多缓存文件，都是微信在使用过程中产生的临时文件。你可以根据自己的需要，勾选不需要的缓存文件，然后点击右下角的“清理”按钮。</p><div class="pgc-img"><img src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/7166eedda5a34026b537aa0787ab2a1d~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=Nvz8xdjiZjHsoLjkehjUHCrMTIA%3D" img_width="1280" img_height="800" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/7166eedda5a34026b537aa0787ab2a1d" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div><p data-track="12">二、关闭不必要的自动下载</p><p data-track="13">1. 关闭小视频和图片的自动下载</p><p data-track="14">回到微信的设置页面，找到“通用”选项（也可能是“数据与通用使用情况”下面的一个选项），然后点击“照片、视频、文件和通话”。在这里，你可以关闭“自动下载”中的“照片”、“视频”和“蜂窝数据”下的“应用首次打开时允许自动加载”。这样，微信就不会自动下载你未点击查看的小视频和图片了。</p><div class="pgc-img"><img src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/0461e265d7b64c70bb6fe1844e8855b7~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=PL2PqkCDI6MPVR2GVEmmzEaXo0M%3D" img_width="800" img_height="1422" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/0461e265d7b64c70bb6fe1844e8855b7" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div><p data-track="15">2. 关闭朋友圈视频自动播放</p><p data-track="16">在“照片、视频、文件和通话”页面，还有一个“视频通话和FaceTime通话”选项。点击进入后，关闭“移动网络下视频自动播放”功能。这样，当你在移动网络下浏览朋友圈时，就不会自动播放小视频了。</p><p data-track="17">三、定期清理聊天记录</p><p data-track="18">虽然简单地点一下“清空”按钮并不能真正释放大量内存，但定期清理聊天记录还是有必要的。你可以选择一些不重要的聊天记录进行删除，或者定期清理一些群聊的聊天记录。</p><div class="pgc-img"><img src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/f7fbaefe394e49af936ef99d0e0a321b~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=ZQ5NJQK3ntPaK%2B3cD5K%2FpwYKV5A%3D" img_width="500" img_height="500" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/f7fbaefe394e49af936ef99d0e0a321b" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div><p data-track="19">四、卸载重装微信</p><p data-track="20">如果你觉得以上方法都不能彻底清理微信，那么你可以尝试卸载微信，然后重新安装。这样，微信的所有数据和缓存都会被清除，但是你的聊天记录也会被清除。所以，在卸载之前，一定要备份好重要的聊天记录。</p><p data-track="21">五、使用第三方工具</p><p data-track="22">如果你觉得手动清理微信太麻烦，或者担心误删重要的聊天记录，那么你可以使用第三方工具来帮助你清理微信。例如，“微信清理君”等清理工具可以帮助你彻底清理微信的内存，而不会误删重要的聊天记录。</p><div class="pgc-img"><img src="https://p6-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/64a989b43de94686b79e0640d2481848~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=2Gx6ok1Npl3xTJh0ONtfu4BWwfA%3D" img_width="400" img_height="711" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/64a989b43de94686b79e0640d2481848" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div><p data-track="23">当然，使用第三方工具时要谨慎，选择那些经过认证的工具，并且在清理前做好数据备份。</p><p data-track="24">以上就是一些清理微信内存的方法，希望对你有所帮助。如果你有其他的好方法，也欢迎在评论区分享哦！</p><div class="pgc-img"><img src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/9ef189c0bc834c608a259d8c12ead187~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=Q2bP3QaVG3I7q4SRKCAQN%2BYx9kA%3D" img_width="500" img_height="889" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/9ef189c0bc834c608a259d8c12ead187" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div><p data-track="25">总之，手机内存是有限的，我们需要学会如何管理和优化内存的使用。对于微信这样的社交软件，我们可以通过深度清理缓存、关闭不必要的自动下载、定期清理聊天记录等方法来释放内存。希望以上方法能对你有所帮助，让你的手机重新焕发活力！</p><div class="pgc-img"><img src="https://p3-sign.toutiaoimg.com/tos-cn-i-6w9my0ksvp/1355f038ba6b42f4815389728ee9500d~noop.image?_iz=58558&amp;from=article.pc_detail&amp;lk3s=953192f4&amp;x-expires=1723512873&amp;x-signature=MxAHOnqEgTlAh664OIk3Yf1GPM8%3D" img_width="1080" img_height="1920" image_type="1" mime_type="image/webp" web_uri="tos-cn-i-6w9my0ksvp/1355f038ba6b42f4815389728ee9500d" class="syl-page-img" style="height: auto;"><p class="pgc-img-caption"></p></div></article>`
	articles := []*entity.Article{
		{
			Title:            "建设一座风力发电机需要花多少钱？一天能发多少电，多久能回本",
			CategoryID:       2,
			OriginAuthor:     "今日头条",
			OriginURL:        "https://www.toutiao.com/article/7385941104601940499/",
			Keywords:         "技术,中小企业健康之道,美好，一直在身边,可再生能源,设计,投资,国创上头条,环境污染",
			Description:      "近年来，我国的风电发电装机容量持续增长，新建风电项目的投建数量位居世界前列，在世界上具备很强的竞争力。",
			ImageURL:         accessor.StaticUrl{FilePath: "/upload/article-image/default-1.jpg"},
			ArticleStatistic: &entity.ArticleStatistic{},
			ArticleContent:   &entity.ArticleContent{Content: content1},
			ArticleTags:      []*entity.ArticleTag{{ID: 5}, {ID: 7}},
		},
		{
			Title:        "微信只清空聊天记录等于没删，教你用正确清理方法能清理好几个G",
			CategoryID:   1,
			OriginAuthor: "今日头条",
			OriginURL:    "https://www.toutiao.com/article/7373860033458733607/",
			Keywords:     "微信,2019科技之光,软件,FaceTime",
			Description:  "微信只清空聊天记录等于没删，教你用正确清理方法能清理好几个G身为一个在互联网上冲浪多年的老鸟，肯定有很多人跟我一样，手机的内存常年告急。而微信，这个社交巨头，往往是内存占用的“大户”。",
			ImageURL:     accessor.StaticUrl{FilePath: "/upload/article-image/default-2.jpg"},
			ArticleStatistic: &entity.ArticleStatistic{
				Views:      rand.UintN(5000),
				Favourites: rand.UintN(500),
			},
			ArticleContent: &entity.ArticleContent{Content: content2},
			ArticleTags:    []*entity.ArticleTag{{ID: 5}, {ID: 3}},
		},
	}

	err := dao.Article.WithContext(c).Create(articles...)

	return err
}
