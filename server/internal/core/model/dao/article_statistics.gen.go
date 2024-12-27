// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"server/internal/core/model/entity"
)

func newArticleStatistic(db *gorm.DB, opts ...gen.DOOption) articleStatistic {
	_articleStatistic := articleStatistic{}

	_articleStatistic.articleStatisticDo.UseDB(db, opts...)
	_articleStatistic.articleStatisticDo.UseModel(&entity.ArticleStatistic{})

	tableName := _articleStatistic.articleStatisticDo.TableName()
	_articleStatistic.ALL = field.NewAsterisk(tableName)
	_articleStatistic.ArticleID = field.NewUint(tableName, "article_id")
	_articleStatistic.UpdatedAt = field.NewTime(tableName, "updated_at")
	_articleStatistic.CreatedAt = field.NewTime(tableName, "created_at")
	_articleStatistic.Views = field.NewUint(tableName, "views")
	_articleStatistic.Favourites = field.NewUint(tableName, "favourites")

	_articleStatistic.fillFieldMap()

	return _articleStatistic
}

// articleStatistic 文章内容
type articleStatistic struct {
	articleStatisticDo articleStatisticDo

	ALL        field.Asterisk
	ArticleID  field.Uint // 文章id
	UpdatedAt  field.Time
	CreatedAt  field.Time
	Views      field.Uint
	Favourites field.Uint

	fieldMap map[string]field.Expr
}

func (a articleStatistic) Table(newTableName string) *articleStatistic {
	a.articleStatisticDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a articleStatistic) As(alias string) *articleStatistic {
	a.articleStatisticDo.DO = *(a.articleStatisticDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *articleStatistic) updateTableName(table string) *articleStatistic {
	a.ALL = field.NewAsterisk(table)
	a.ArticleID = field.NewUint(table, "article_id")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.Views = field.NewUint(table, "views")
	a.Favourites = field.NewUint(table, "favourites")

	a.fillFieldMap()

	return a
}

func (a *articleStatistic) WithContext(ctx context.Context) IArticleStatisticDo {
	return a.articleStatisticDo.WithContext(ctx)
}

func (a articleStatistic) TableName() string { return a.articleStatisticDo.TableName() }

func (a articleStatistic) Alias() string { return a.articleStatisticDo.Alias() }

func (a articleStatistic) Columns(cols ...field.Expr) gen.Columns {
	return a.articleStatisticDo.Columns(cols...)
}

func (a *articleStatistic) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *articleStatistic) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 5)
	a.fieldMap["article_id"] = a.ArticleID
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["views"] = a.Views
	a.fieldMap["favourites"] = a.Favourites
}

func (a articleStatistic) clone(db *gorm.DB) articleStatistic {
	a.articleStatisticDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a articleStatistic) replaceDB(db *gorm.DB) articleStatistic {
	a.articleStatisticDo.ReplaceDB(db)
	return a
}

type articleStatisticDo struct{ gen.DO }

type IArticleStatisticDo interface {
	gen.SubQuery
	Debug() IArticleStatisticDo
	WithContext(ctx context.Context) IArticleStatisticDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IArticleStatisticDo
	WriteDB() IArticleStatisticDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IArticleStatisticDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IArticleStatisticDo
	Not(conds ...gen.Condition) IArticleStatisticDo
	Or(conds ...gen.Condition) IArticleStatisticDo
	Select(conds ...field.Expr) IArticleStatisticDo
	Where(conds ...gen.Condition) IArticleStatisticDo
	Order(conds ...field.Expr) IArticleStatisticDo
	Distinct(cols ...field.Expr) IArticleStatisticDo
	Omit(cols ...field.Expr) IArticleStatisticDo
	Join(table schema.Tabler, on ...field.Expr) IArticleStatisticDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IArticleStatisticDo
	RightJoin(table schema.Tabler, on ...field.Expr) IArticleStatisticDo
	Group(cols ...field.Expr) IArticleStatisticDo
	Having(conds ...gen.Condition) IArticleStatisticDo
	Limit(limit int) IArticleStatisticDo
	Offset(offset int) IArticleStatisticDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleStatisticDo
	Unscoped() IArticleStatisticDo
	Create(values ...*entity.ArticleStatistic) error
	CreateInBatches(values []*entity.ArticleStatistic, batchSize int) error
	Save(values ...*entity.ArticleStatistic) error
	First() (*entity.ArticleStatistic, error)
	Take() (*entity.ArticleStatistic, error)
	Last() (*entity.ArticleStatistic, error)
	Find() ([]*entity.ArticleStatistic, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.ArticleStatistic, err error)
	FindInBatches(result *[]*entity.ArticleStatistic, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.ArticleStatistic) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IArticleStatisticDo
	Assign(attrs ...field.AssignExpr) IArticleStatisticDo
	Joins(fields ...field.RelationField) IArticleStatisticDo
	Preload(fields ...field.RelationField) IArticleStatisticDo
	FirstOrInit() (*entity.ArticleStatistic, error)
	FirstOrCreate() (*entity.ArticleStatistic, error)
	FindByPage(offset int, limit int) (result []*entity.ArticleStatistic, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IArticleStatisticDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a articleStatisticDo) Debug() IArticleStatisticDo {
	return a.withDO(a.DO.Debug())
}

func (a articleStatisticDo) WithContext(ctx context.Context) IArticleStatisticDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a articleStatisticDo) ReadDB() IArticleStatisticDo {
	return a.Clauses(dbresolver.Read)
}

func (a articleStatisticDo) WriteDB() IArticleStatisticDo {
	return a.Clauses(dbresolver.Write)
}

func (a articleStatisticDo) Session(config *gorm.Session) IArticleStatisticDo {
	return a.withDO(a.DO.Session(config))
}

func (a articleStatisticDo) Clauses(conds ...clause.Expression) IArticleStatisticDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a articleStatisticDo) Returning(value interface{}, columns ...string) IArticleStatisticDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a articleStatisticDo) Not(conds ...gen.Condition) IArticleStatisticDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a articleStatisticDo) Or(conds ...gen.Condition) IArticleStatisticDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a articleStatisticDo) Select(conds ...field.Expr) IArticleStatisticDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a articleStatisticDo) Where(conds ...gen.Condition) IArticleStatisticDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a articleStatisticDo) Order(conds ...field.Expr) IArticleStatisticDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a articleStatisticDo) Distinct(cols ...field.Expr) IArticleStatisticDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a articleStatisticDo) Omit(cols ...field.Expr) IArticleStatisticDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a articleStatisticDo) Join(table schema.Tabler, on ...field.Expr) IArticleStatisticDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a articleStatisticDo) LeftJoin(table schema.Tabler, on ...field.Expr) IArticleStatisticDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a articleStatisticDo) RightJoin(table schema.Tabler, on ...field.Expr) IArticleStatisticDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a articleStatisticDo) Group(cols ...field.Expr) IArticleStatisticDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a articleStatisticDo) Having(conds ...gen.Condition) IArticleStatisticDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a articleStatisticDo) Limit(limit int) IArticleStatisticDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a articleStatisticDo) Offset(offset int) IArticleStatisticDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a articleStatisticDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IArticleStatisticDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a articleStatisticDo) Unscoped() IArticleStatisticDo {
	return a.withDO(a.DO.Unscoped())
}

func (a articleStatisticDo) Create(values ...*entity.ArticleStatistic) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a articleStatisticDo) CreateInBatches(values []*entity.ArticleStatistic, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a articleStatisticDo) Save(values ...*entity.ArticleStatistic) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a articleStatisticDo) First() (*entity.ArticleStatistic, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ArticleStatistic), nil
	}
}

func (a articleStatisticDo) Take() (*entity.ArticleStatistic, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ArticleStatistic), nil
	}
}

func (a articleStatisticDo) Last() (*entity.ArticleStatistic, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ArticleStatistic), nil
	}
}

func (a articleStatisticDo) Find() ([]*entity.ArticleStatistic, error) {
	result, err := a.DO.Find()
	return result.([]*entity.ArticleStatistic), err
}

func (a articleStatisticDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.ArticleStatistic, err error) {
	buf := make([]*entity.ArticleStatistic, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a articleStatisticDo) FindInBatches(result *[]*entity.ArticleStatistic, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a articleStatisticDo) Attrs(attrs ...field.AssignExpr) IArticleStatisticDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a articleStatisticDo) Assign(attrs ...field.AssignExpr) IArticleStatisticDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a articleStatisticDo) Joins(fields ...field.RelationField) IArticleStatisticDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a articleStatisticDo) Preload(fields ...field.RelationField) IArticleStatisticDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a articleStatisticDo) FirstOrInit() (*entity.ArticleStatistic, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ArticleStatistic), nil
	}
}

func (a articleStatisticDo) FirstOrCreate() (*entity.ArticleStatistic, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.ArticleStatistic), nil
	}
}

func (a articleStatisticDo) FindByPage(offset int, limit int) (result []*entity.ArticleStatistic, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a articleStatisticDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a articleStatisticDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a articleStatisticDo) Delete(models ...*entity.ArticleStatistic) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *articleStatisticDo) withDO(do gen.Dao) *articleStatisticDo {
	a.DO = *do.(*gen.DO)
	return a
}
