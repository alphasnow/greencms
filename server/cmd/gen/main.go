// @author AlphaSnow

package main

import (
	"golang.org/x/tools/go/packages"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var dbDSN = "root:root@(127.0.0.1:3306)/oneclick?charset=utf8mb4&parseTime=True&loc=Local"

// 目录结构参考
// https://goframe.org/pages/viewpage.action?pageId=30740166
var genOutPath = "./internal/core/model/dao"

// var gemYamlFile = "./scripts/gen/gen.yaml"
var accessorPath = "./internal/core/model/accessor"

func main() {

	g := connection()

	withOpts(g)
	generateAll(g)

	generateModelRelate(g)
	//generateModelRelateByXml(g)

	modifyField(g)

	g.Execute()

}

func generateModelRelate(g *gen.Generator) {
	web_metas := g.GenerateModel("web_metas")
	g.ApplyBasic(web_metas)

	web_banners := g.GenerateModel("web_banners",
		gen.FieldType("image_url", "accessor.StaticUrl"),
	)
	g.ApplyBasic(web_banners)

	admin_users := g.GenerateModel("admin_users",
		gen.FieldType("avatar_url", "accessor.StaticUrl"),
	)
	g.ApplyBasic(admin_users)

	article_categories := g.GenerateModel("article_categories",
		gen.FieldType("image_url", "accessor.StaticUrl"),
	)
	g.ApplyBasic(article_categories)
	article_contents := g.GenerateModel("article_contents")
	g.ApplyBasic(article_contents)
	article_tag_relates := g.GenerateModel("article_tag_relates")
	g.ApplyBasic(article_tag_relates)
	article_tags := g.GenerateModel("article_tags")
	g.ApplyBasic(article_tags)
	article_statistics := g.GenerateModel("article_statistics")
	g.ApplyBasic(article_statistics)

	opts := []gen.ModelOpt{
		gen.FieldType("image_url", "accessor.StaticUrl"),
		// gen.FieldJSONTag("deleted_at", "-"),
		gen.FieldRelate(field.BelongsTo, article_categories.ModelStructName, article_categories, &field.RelateConfig{
			RelatePointer: true,
			GORMTag:       field.GormTag{"foreignKey": []string{"CategoryID"}},
			JSONTag:       `article_category,omitempty`,
		}),
		gen.FieldRelate(field.HasOne, article_contents.ModelStructName, article_contents, &field.RelateConfig{
			RelatePointer: true,
			GORMTag:       field.GormTag{"foreignKey": []string{"ArticleID"}},
			JSONTag:       `article_content,omitempty`,
		}),
		gen.FieldRelate(field.HasOne, article_statistics.ModelStructName, article_statistics, &field.RelateConfig{
			RelatePointer: true,
			GORMTag:       field.GormTag{"foreignKey": []string{"ArticleID"}},
			JSONTag:       `article_statistic,omitempty`,
		}),
		gen.FieldRelate(field.Many2Many, "ArticleTags", article_tags, &field.RelateConfig{
			RelatePointer:      false,
			RelateSlicePointer: true,
			JSONTag:            `article_tags,omitempty`,
			GORMTag: field.GormTag{
				"Many2many":      []string{"article_tag_relates"},
				"foreignKey":     []string{"ID"},
				"references":     []string{"ID"},
				"joinForeignKey": []string{"ArticleID"},
				"joinReferences": []string{"TagID"},
			},
		}),
	}
	article := g.GenerateModel("articles", opts...)
	g.ApplyBasic(article)
}

func generateModelRelateByXml(g *gen.Generator) {
	//yamlgen 可以使用yaml配置关系
	//不支持多态 不支持自定义关联字段名称

	//xyamlgen.NewYamlGenerator(gemYamlFile).UseGormGenerator(g).Generate()
}

func modifyField(g *gen.Generator) {
	//// 添加自定义参数
	//webAttachment := g.GenerateModel("web_attachments",
	//	gen.FieldNew("FileURL", "string", field.Tag{
	//		"gorm": "-",
	//		"json": "file_url,omitempty",
	//	}),
	//)
	//g.ApplyBasic(webAttachment)
	//
	//articleCategory := g.GenerateModel("article_categories",
	//	gen.FieldType("cover_url", "accessor.StaticUrl"),
	//)
	//g.ApplyBasic(articleCategory)
}

func connection() *gen.Generator {
	c := gen.Config{
		OutPath:      genOutPath,
		ModelPkgPath: "entity",

		//Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
		////if you want the nullable field generation property to be pointer type, set FieldNullable true
		//FieldNullable: true,
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		// if you want generate field with unsigned integer type, set FieldSignable true
		//FieldSignable: true,
		//if you want to generate index tags from database, set FieldWithIndexTag true
		//FieldWithIndexTag: true,
		// sqlite 会受影响需要关闭 FieldWithTypeTag
		//if you want to generate type tags from database, set FieldWithTypeTag true
		//FieldWithTypeTag: true,
		//if you need unit tests for query code, set WithUnitTest true
		//WithUnitTest: true,
	}
	g := gen.NewGenerator(c)

	// 连接数据库 创建模型
	db, _ := gorm.Open(mysql.Open(dbDSN))
	g.UseDB(db)

	// 不连接数据库 通过本地的sql文件创建模型
	// https://gorm.io/gen/database_to_structs.html#Generate-From-Sql
	//gormdb, _ := gorm.Open(rawsql.New(rawsql.Config{
	//	//SQL:      rawsql,                      //create table sql
	//	FilePath: []string{
	//		"./storage/data/database.db", // create table sql file directory
	//	},
	//}))
	//g.UseDB(gormdb)

	return g
}

func withOpts(g *gen.Generator) {
	// [MAP]
	// https://gorm.io/gen/database_to_structs.html#Data-Mapping
	dMap := map[string]func(gorm.ColumnType) (dataType string){
		// int / unit32
		"int": func(columnType gorm.ColumnType) (dataType string) {
			//if n, ok := columnType.Nullable(); ok && n {
			//	return "*uint"
			//}
			return "uint"
		},
		// bigint / uint64
		"bigint": func(columnType gorm.ColumnType) (dataType string) {
			//if n, ok := columnType.Nullable(); ok && n {
			//	return "*uint"
			//}
			return "uint"
		},
		"timestamp": func(columnType gorm.ColumnType) (dataType string) {
			if columnType.Name() == "created_at" || columnType.Name() == "updated_at" {
				return "time.Time"
			} else if columnType.Name() == "deleted_at" {
				return "gorm.DeletedAt"
			}

			if n, ok := columnType.Nullable(); ok && n {
				return "*time.Time"
			} else {
				return "time.Time"
			}

		},
		//
		//// bool mapping
		//"tinyint": func(columnType gorm.ColumnType) (dataType string) {
		//	ct, _ := columnType.ColumnType()
		//	if strings.HasPrefix(ct, "tinyint(1)") {
		//		return "bool"
		//	}
		//	return "byte"
		//},
	}
	g.WithDataTypeMap(dMap)

	// [OPT]
	fieldOpts := []gen.ModelOpt{
		gen.FieldJSONTag("deleted_at", "-"),
		gen.FieldJSONTag("password", "-"),
		gen.FieldGORMTag("created_at", func(tag field.GormTag) field.GormTag {
			tag.Append("<-", "create")
			return tag
		}),
		//gen.FieldNew("FileUrl", "string", field.Tag{
		//	"gorm": "-",
		//	"json": "file_url,omitempty",
		//}),
	}
	g.WithOpts(fieldOpts...)

	// [PKG]
	pks, _ := packages.Load(&packages.Config{
		Mode: packages.NeedName,
		Dir:  accessorPath,
	})
	g.WithImportPkgPath(pks[0].PkgPath)
}

func generateAll(g *gen.Generator) {
	// 根据已有模型生成dao
	// g.ApplyBasic(model.Article{})
	// 生成所有表的 model,
	// Error 会引起生成的dao没有关联字段
	// g.ApplyBasic(g.GenerateAllTable()...)
}
