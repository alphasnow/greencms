package accessor

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"server/pkg/g"
	"strings"
)

// StaticUrl
// 自定义类型 https://gorm.io/docs/data_types.html
// 也可以考虑用AfterFind解决 https://gorm.io/docs/hooks.html#Querying-an-object
type StaticUrl struct {
	FilePath string
	FileUrl  string
}

// Scan
// value into Jsonb, implements sql.Scanner interface
// 从数据库读取后 添加域名
func (s *StaticUrl) Scan(value interface{}) error {
	if value == nil {
		*s = StaticUrl{}
		return nil
	}

	v, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan StaticUrl: %v", value)
	}

	full := s.pathToUrl(v)
	*s = StaticUrl{FilePath: v, FileUrl: full}
	return nil
}

// Value
// return json value, implement driver.Valuer interface
// 存储数据库前 去除域名
func (s StaticUrl) Value() (driver.Value, error) {
	return s.FilePath, nil
}

func (StaticUrl) GormDataType() string {
	return "string"
}

func (StaticUrl) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "text"
	}
	return "varchar"
}

func (s *StaticUrl) String() string {
	return s.FileUrl
}

func (s StaticUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.FileUrl)
}

func (s *StaticUrl) UnmarshalJSON(data []byte) error {
	// ignore null
	if len(data) == 0 || string(data) == "null" {
		*s = StaticUrl{}
		return nil
	}
	var str string
	err := json.Unmarshal(data, &str)
	if err != nil {
		return fmt.Errorf("failed to UnmarshalJSON StaticUrl: %s", err)
	}

	lp := s.urlToPath(str)
	*s = StaticUrl{FilePath: lp, FileUrl: str}
	return nil
}

func (s *StaticUrl) urlToPath(val string) string {
	prefix := g.Url("/")
	// 若是远程地址则不处理
	if strings.HasPrefix(val, prefix) == false {
		return val
	}
	str := strings.TrimPrefix(val, prefix)
	return str
}
func (s *StaticUrl) pathToUrl(val string) string {
	// 若是远程地址则不处理
	if strings.HasPrefix(val, "http://") || strings.HasPrefix(val, "https://") {
		return val
	}
	return g.Url(val)
}
