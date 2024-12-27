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

func newWebMeta(db *gorm.DB, opts ...gen.DOOption) webMeta {
	_webMeta := webMeta{}

	_webMeta.webMetaDo.UseDB(db, opts...)
	_webMeta.webMetaDo.UseModel(&entity.WebMeta{})

	tableName := _webMeta.webMetaDo.TableName()
	_webMeta.ALL = field.NewAsterisk(tableName)
	_webMeta.ID = field.NewUint(tableName, "id")
	_webMeta.CreatedAt = field.NewTime(tableName, "created_at")
	_webMeta.UpdatedAt = field.NewTime(tableName, "updated_at")
	_webMeta.DeletedAt = field.NewField(tableName, "deleted_at")
	_webMeta.MetaKey = field.NewString(tableName, "meta_key")
	_webMeta.MetaValue = field.NewString(tableName, "meta_value")
	_webMeta.MetaGroup = field.NewString(tableName, "meta_group")
	_webMeta.MetaName = field.NewString(tableName, "meta_name")
	_webMeta.Remark = field.NewString(tableName, "remark")

	_webMeta.fillFieldMap()

	return _webMeta
}

// webMeta 元数据
type webMeta struct {
	webMetaDo webMetaDo

	ALL       field.Asterisk
	ID        field.Uint // 主键
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	MetaKey   field.String
	MetaValue field.String
	MetaGroup field.String
	MetaName  field.String
	Remark    field.String

	fieldMap map[string]field.Expr
}

func (w webMeta) Table(newTableName string) *webMeta {
	w.webMetaDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w webMeta) As(alias string) *webMeta {
	w.webMetaDo.DO = *(w.webMetaDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *webMeta) updateTableName(table string) *webMeta {
	w.ALL = field.NewAsterisk(table)
	w.ID = field.NewUint(table, "id")
	w.CreatedAt = field.NewTime(table, "created_at")
	w.UpdatedAt = field.NewTime(table, "updated_at")
	w.DeletedAt = field.NewField(table, "deleted_at")
	w.MetaKey = field.NewString(table, "meta_key")
	w.MetaValue = field.NewString(table, "meta_value")
	w.MetaGroup = field.NewString(table, "meta_group")
	w.MetaName = field.NewString(table, "meta_name")
	w.Remark = field.NewString(table, "remark")

	w.fillFieldMap()

	return w
}

func (w *webMeta) WithContext(ctx context.Context) IWebMetaDo { return w.webMetaDo.WithContext(ctx) }

func (w webMeta) TableName() string { return w.webMetaDo.TableName() }

func (w webMeta) Alias() string { return w.webMetaDo.Alias() }

func (w webMeta) Columns(cols ...field.Expr) gen.Columns { return w.webMetaDo.Columns(cols...) }

func (w *webMeta) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *webMeta) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 9)
	w.fieldMap["id"] = w.ID
	w.fieldMap["created_at"] = w.CreatedAt
	w.fieldMap["updated_at"] = w.UpdatedAt
	w.fieldMap["deleted_at"] = w.DeletedAt
	w.fieldMap["meta_key"] = w.MetaKey
	w.fieldMap["meta_value"] = w.MetaValue
	w.fieldMap["meta_group"] = w.MetaGroup
	w.fieldMap["meta_name"] = w.MetaName
	w.fieldMap["remark"] = w.Remark
}

func (w webMeta) clone(db *gorm.DB) webMeta {
	w.webMetaDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w webMeta) replaceDB(db *gorm.DB) webMeta {
	w.webMetaDo.ReplaceDB(db)
	return w
}

type webMetaDo struct{ gen.DO }

type IWebMetaDo interface {
	gen.SubQuery
	Debug() IWebMetaDo
	WithContext(ctx context.Context) IWebMetaDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWebMetaDo
	WriteDB() IWebMetaDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWebMetaDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWebMetaDo
	Not(conds ...gen.Condition) IWebMetaDo
	Or(conds ...gen.Condition) IWebMetaDo
	Select(conds ...field.Expr) IWebMetaDo
	Where(conds ...gen.Condition) IWebMetaDo
	Order(conds ...field.Expr) IWebMetaDo
	Distinct(cols ...field.Expr) IWebMetaDo
	Omit(cols ...field.Expr) IWebMetaDo
	Join(table schema.Tabler, on ...field.Expr) IWebMetaDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWebMetaDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWebMetaDo
	Group(cols ...field.Expr) IWebMetaDo
	Having(conds ...gen.Condition) IWebMetaDo
	Limit(limit int) IWebMetaDo
	Offset(offset int) IWebMetaDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWebMetaDo
	Unscoped() IWebMetaDo
	Create(values ...*entity.WebMeta) error
	CreateInBatches(values []*entity.WebMeta, batchSize int) error
	Save(values ...*entity.WebMeta) error
	First() (*entity.WebMeta, error)
	Take() (*entity.WebMeta, error)
	Last() (*entity.WebMeta, error)
	Find() ([]*entity.WebMeta, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.WebMeta, err error)
	FindInBatches(result *[]*entity.WebMeta, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.WebMeta) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWebMetaDo
	Assign(attrs ...field.AssignExpr) IWebMetaDo
	Joins(fields ...field.RelationField) IWebMetaDo
	Preload(fields ...field.RelationField) IWebMetaDo
	FirstOrInit() (*entity.WebMeta, error)
	FirstOrCreate() (*entity.WebMeta, error)
	FindByPage(offset int, limit int) (result []*entity.WebMeta, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWebMetaDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w webMetaDo) Debug() IWebMetaDo {
	return w.withDO(w.DO.Debug())
}

func (w webMetaDo) WithContext(ctx context.Context) IWebMetaDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w webMetaDo) ReadDB() IWebMetaDo {
	return w.Clauses(dbresolver.Read)
}

func (w webMetaDo) WriteDB() IWebMetaDo {
	return w.Clauses(dbresolver.Write)
}

func (w webMetaDo) Session(config *gorm.Session) IWebMetaDo {
	return w.withDO(w.DO.Session(config))
}

func (w webMetaDo) Clauses(conds ...clause.Expression) IWebMetaDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w webMetaDo) Returning(value interface{}, columns ...string) IWebMetaDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w webMetaDo) Not(conds ...gen.Condition) IWebMetaDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w webMetaDo) Or(conds ...gen.Condition) IWebMetaDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w webMetaDo) Select(conds ...field.Expr) IWebMetaDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w webMetaDo) Where(conds ...gen.Condition) IWebMetaDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w webMetaDo) Order(conds ...field.Expr) IWebMetaDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w webMetaDo) Distinct(cols ...field.Expr) IWebMetaDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w webMetaDo) Omit(cols ...field.Expr) IWebMetaDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w webMetaDo) Join(table schema.Tabler, on ...field.Expr) IWebMetaDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w webMetaDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWebMetaDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w webMetaDo) RightJoin(table schema.Tabler, on ...field.Expr) IWebMetaDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w webMetaDo) Group(cols ...field.Expr) IWebMetaDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w webMetaDo) Having(conds ...gen.Condition) IWebMetaDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w webMetaDo) Limit(limit int) IWebMetaDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w webMetaDo) Offset(offset int) IWebMetaDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w webMetaDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWebMetaDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w webMetaDo) Unscoped() IWebMetaDo {
	return w.withDO(w.DO.Unscoped())
}

func (w webMetaDo) Create(values ...*entity.WebMeta) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w webMetaDo) CreateInBatches(values []*entity.WebMeta, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w webMetaDo) Save(values ...*entity.WebMeta) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w webMetaDo) First() (*entity.WebMeta, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebMeta), nil
	}
}

func (w webMetaDo) Take() (*entity.WebMeta, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebMeta), nil
	}
}

func (w webMetaDo) Last() (*entity.WebMeta, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebMeta), nil
	}
}

func (w webMetaDo) Find() ([]*entity.WebMeta, error) {
	result, err := w.DO.Find()
	return result.([]*entity.WebMeta), err
}

func (w webMetaDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.WebMeta, err error) {
	buf := make([]*entity.WebMeta, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w webMetaDo) FindInBatches(result *[]*entity.WebMeta, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w webMetaDo) Attrs(attrs ...field.AssignExpr) IWebMetaDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w webMetaDo) Assign(attrs ...field.AssignExpr) IWebMetaDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w webMetaDo) Joins(fields ...field.RelationField) IWebMetaDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w webMetaDo) Preload(fields ...field.RelationField) IWebMetaDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w webMetaDo) FirstOrInit() (*entity.WebMeta, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebMeta), nil
	}
}

func (w webMetaDo) FirstOrCreate() (*entity.WebMeta, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.WebMeta), nil
	}
}

func (w webMetaDo) FindByPage(offset int, limit int) (result []*entity.WebMeta, count int64, err error) {
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

func (w webMetaDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w webMetaDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w webMetaDo) Delete(models ...*entity.WebMeta) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *webMetaDo) withDO(do gen.Dao) *webMetaDo {
	w.DO = *do.(*gen.DO)
	return w
}