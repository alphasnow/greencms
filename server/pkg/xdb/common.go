package xdb

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(username, password, address, dbname, options string) (*gorm.DB, error) {
	dsn := toSDN(username, password, address, dbname, options)
	conn := mysql.Open(dsn)
	db, err := gorm.Open(conn, &gorm.Config{
		// you can disable it during initialization if it is not required, you will gain about 30%+ performance improvement after that
		// https://gorm.io/docs/transactions.html
		SkipDefaultTransaction: true,
	})
	return db, err
}

func toSDN(username, password, address, dbname, options string) string {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", username, password, address, dbname, options)
}

func ping(db *gorm.DB) error {
	d, err := db.DB()
	if err != nil {
		return err
	}
	return d.Ping()
}

func toDSN(username, password, address, options string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/?%s", username, password, address, options)
}

func createDbname(dsn string, dbname string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET = `utf8mb4`;", dbname)
	_, err = db.Exec(query)
	return err
}
