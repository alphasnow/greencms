package g

import (
	"gorm.io/gorm"
	"server/pkg/xdb"
)

func DB() (db *gorm.DB) {
	var err error
	if globalContainer.Has("db") {
		db = globalContainer.Get("db").(*gorm.DB)
	} else {
		db, err = xdb.ProvideDB(Config(), Log("server"))
		if err != nil {
			panic(err)
		}
		globalContainer.Set("db", db)
	}
	return db
}
