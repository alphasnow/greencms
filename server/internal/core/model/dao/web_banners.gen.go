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

func newWebBanner(db *gorm.DB, opts ...gen.DOOption) webBanner {
	_webBanner := webBanner{}

	_webBanner.webBannerDo.UseDB(db, opts...)
	_webBanner.webBannerDo.UseModel(&entity.WebBanner{})

	tableName := _webBanner.webBannerDo.TableName()
	_webBanner.ALL = field.NewAsterisk(tableName)
	_webBanner.ID = field.NewUint(tableName, "id")
	_webBanner.CreatedAt = field.NewTime(tableName, "created_at")
	_webBanner.UpdatedAt = field.NewTime(tableName, "updated_at")
	_webBanner.DeletedAt = field.NewField(tableName, "deleted_at")
	_webBanner.ImageURL = field.NewField(tableName, "image_url")
	_webBanner.RedirectURL = field.NewString(tableName, "redirect_url")
	_webBanner.BannerGroup = field.NewString(tableName, "banner_group")
	_webBanner.Sort = field.NewInt32(tableName, "sort")
	_webBanner.Remark = field.NewString(tableName, "remark")
	_webBanner.Title = field.NewString(tableName, "title")
	_webBanner.Description = field.NewString(tableName, "description")

	_webBanner.fillFieldMap()

	return _webBanner
}

// webBanner 轮播图
type webBanner struct {
	webBannerDo webBannerDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	ImageURL    field.Field
	RedirectURL field.String
	BannerGroup field.String
	Sort        field.Int32
	Remark      field.String
	Title       field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (w webBanner) Table(newTableName string) *webBanner {
	w.webBannerDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w webBanner) As(alias string) *webBanner {
	w.webBannerDo.DO = *(w.webBannerDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *webBanner) updateTableName(table string) *webBanner {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.ImageURL = field.NewField(table, "image_url")
	w.RedirectURL = field.NewString(table, "redirect_url")
	w.BannerGroup = field.NewString(table, "banner_group")
	w.Sort = field.NewInt32(table, "sort")
	w.Remark = field.NewString(table, "remark")
	w.Title = field.NewString(table, "title")
	w.Description = field.NewString(table, "description")

	w.fillFieldMap()

	return w
}

func (w *webBanner) WithContext(ctx context.Context) IWebBannerDo {
	return w.webBannerDo.WithContext(ctx)
}

func (w webBanner) TableName() string { return w.webBannerDo.TableName() }

func (w webBanner) Alias() string { return w.webBannerDo.Alias() }

func (w webBanner) Columns(cols ...field.Expr) gen.Columns { return w.webBannerDo.Columns(cols...) }

func (w *webBanner) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *webBanner) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 11)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["image_url"] = w.ImageURL
	w.fieldMap["redirect_url"] = w.RedirectURL
	w.fieldMap["banner_group"] = w.BannerGroup
	w.fieldMap["sort"] = w.Sort
	w.fieldMap["remark"] = w.Remark
	w.fieldMap["title"] = w.Title
	w.fieldMap["description"] = w.Description
}

func (w webBanner) clone(db *gorm.DB) webBanner {
	w.webBannerDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w webBanner) replaceDB(db *gorm.DB) webBanner {
	w.webBannerDo.ReplaceDB(db)
	return w
}

type webBannerDo struct{ gen.DO }

type IWebBannerDo interface {
	gen.SubQuery
	Debug() IWebBannerDo
	WithContext(ctx context.Context) IWebBannerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWebBannerDo
	WriteDB() IWebBannerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWebBannerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWebBannerDo
	Not(conds ...gen.Condition) IWebBannerDo
	Or(conds ...gen.Condition) IWebBannerDo
	Select(conds ...field.Expr) IWebBannerDo
	Where(conds ...gen.Condition) IWebBannerDo
	Order(conds ...field.Expr) IWebBannerDo
	Distinct(cols ...field.Expr) IWebBannerDo
	Omit(cols ...field.Expr) IWebBannerDo
	Join(table schema.Tabler, on ...field.Expr) IWebBannerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWebBannerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWebBannerDo
	Group(cols ...field.Expr) IWebBannerDo
	Having(conds ...gen.Condition) IWebBannerDo
	Limit(limit int) IWebBannerDo
	Offset(offset int) IWebBannerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWebBannerDo
	Unscoped() IWebBannerDo
	Create(values ...*entity.WebBanner) error
	CreateInBatches(values []*entity.WebBanner, batchSize int) error
	Save(values ...*entity.WebBanner) error
	First() (*entity.WebBanner, error)
	Take() (*entity.WebBanner, error)
	Last() (*entity.WebBanner, error)
	Find() ([]*entity.WebBanner, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.WebBanner, err error)
	FindInBatches(result *[]*entity.WebBanner, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.WebBanner) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWebBannerDo
	Assign(attrs ...field.AssignExpr) IWebBannerDo
	Joins(fields ...field.RelationField) IWebBannerDo
	Preload(fields ...field.RelationField) IWebBannerDo
	FirstOrInit() (*entity.WebBanner, error)
	FirstOrCreate() (*entity.WebBanner, error)
	FindByPage(offset int, limit int) (result []*entity.WebBanner, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWebBannerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w webBannerDo) Debug() IWebBannerDo {
	return w.withDO(w.DO.Debug())
}

func (w webBannerDo) WithContext(ctx context.Context) IWebBannerDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w webBannerDo) ReadDB() IWebBannerDo {
	return w.Clauses(dbresolver.Read)
}

func (w webBannerDo) WriteDB() IWebBannerDo {
	return w.Clauses(dbresolver.Write)
}

func (w webBannerDo) Session(config *gorm.Session) IWebBannerDo {
	return w.withDO(w.DO.Session(config))
}

func (w webBannerDo) Clauses(conds ...clause.Expression) IWebBannerDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w webBannerDo) Returning(value interface{}, columns ...string) IWebBannerDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w webBannerDo) Not(conds ...gen.Condition) IWebBannerDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w webBannerDo) Or(conds ...gen.Condition) IWebBannerDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w webBannerDo) Select(conds ...field.Expr) IWebBannerDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w webBannerDo) Where(conds ...gen.Condition) IWebBannerDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w webBannerDo) Order(conds ...field.Expr) IWebBannerDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w webBannerDo) Distinct(cols ...field.Expr) IWebBannerDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w webBannerDo) Omit(cols ...field.Expr) IWebBannerDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w webBannerDo) Join(table schema.Tabler, on ...field.Expr) IWebBannerDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w webBannerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWebBannerDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w webBannerDo) RightJoin(table schema.Tabler, on ...field.Expr) IWebBannerDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w webBannerDo) Group(cols ...field.Expr) IWebBannerDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w webBannerDo) Having(conds ...gen.Condition) IWebBannerDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w webBannerDo) Limit(limit int) IWebBannerDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w webBannerDo) Offset(offset int) IWebBannerDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w webBannerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWebBannerDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w webBannerDo) Unscoped() IWebBannerDo {
	return w.withDO(w.DO.Unscoped())
}

func (w webBannerDo) Create(values ...*entity.WebBanner) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w webBannerDo) CreateInBatches(values []*entity.WebBanner, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w webBannerDo) Save(values ...*entity.WebBanner) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w webBannerDo) First() (*entity.WebBanner, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebBanner), nil
	}
}

func (w webBannerDo) Take() (*entity.WebBanner, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebBanner), nil
	}
}

func (w webBannerDo) Last() (*entity.WebBanner, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebBanner), nil
	}
}

func (w webBannerDo) Find() ([]*entity.WebBanner, error) {
	result, err := w.DO.Find()
	return result.([]*entity.WebBanner), err
}

func (w webBannerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.WebBanner, err error) {
	buf := make([]*entity.WebBanner, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w webBannerDo) FindInBatches(result *[]*entity.WebBanner, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w webBannerDo) Attrs(attrs ...field.AssignExpr) IWebBannerDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w webBannerDo) Assign(attrs ...field.AssignExpr) IWebBannerDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w webBannerDo) Joins(fields ...field.RelationField) IWebBannerDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w webBannerDo) Preload(fields ...field.RelationField) IWebBannerDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w webBannerDo) FirstOrInit() (*entity.WebBanner, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebBanner), nil
	}
}

func (w webBannerDo) FirstOrCreate() (*entity.WebBanner, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebBanner), nil
	}
}

func (w webBannerDo) FindByPage(offset int, limit int) (result []*entity.WebBanner, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w webBannerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w webBannerDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w webBannerDo) Delete(models ...*entity.WebBanner) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *webBannerDo) withDO(do gen.Dao) *webBannerDo {
	w.DO = *do.(*gen.DO)
	return w
}