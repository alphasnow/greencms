// @author AlphaSnow

package xdb

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
	"time"
)

func ProvideDB(conf *viper.Viper, logger *zap.Logger) (db *gorm.DB, err error) {
	//if conf.GetString("database.default") != "mysql" {
	//	panic("database only support mysql")
	//}

	//logger := newGormLogger(conf, l)
	var conn gorm.Dialector
	def := conf.GetString("database.default")
	switch def {
	case "mysql":
		conn = openMySQL(
			conf.GetString("database.connections.mysql.username"),
			conf.GetString("database.connections.mysql.password"),
			conf.GetString("database.connections.mysql.addr"),
			conf.GetString("database.connections.mysql.db"),
			conf.GetString("database.connections.mysql.options"),
		)
		break
	case "sqlite":
		conn = openSqlLite(
			conf.GetString("database.connections.sqlite.filename"),
		)
		break
	default:
		err = fmt.Errorf("invalid value for db.default: %s", conf.GetString("database.default"))
	}

	// logger
	//logger2 := zapgorm2.New(logger)
	//logger2.SetAsDefault()
	//logger2.LogLevel = gormlogger.Warn
	//logger2.SlowThreshold = 500 * time.Millisecond
	logger2 := zapgorm2.Logger{
		ZapLogger:                 logger,
		LogLevel:                  gormlogger.Warn,
		SlowThreshold:             200 * time.Millisecond,
		SkipCallerLookup:          true,
		IgnoreRecordNotFoundError: true,
		Context:                   nil,
	}

	// db
	db, err = gorm.Open(conn, &gorm.Config{
		Logger:                 logger2,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return nil, err
	}

	debug := conf.GetBool("app.debug")
	if debug {
		db = db.Debug()
	}

	//cleanup = func() {
	//	var sd *sql.DB
	//	if sd, err = db.DB(); err == nil {
	//		if err = sd.Close(); err != nil {
	//			log.Printf("close db error: %v", err)
	//		}
	//	}
	//}

	//SetDefault(db)

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 配置连接池
	// 默认 2
	sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetConnMaxIdleTime(30 * time.Minute)
	// 默认 不限制, 限制最大连接数, 控制连接占用资源
	// sqlDB.SetMaxOpenConns(100)
	// 默认 0
	// sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func openMySQL(username, password, address, dbname, options string) gorm.Dialector {
	dsn := toSDN(
		username, password, address, dbname, options,
	)
	conn := mysql.Open(dsn)
	return conn
}

func openSqlLite(filename string) gorm.Dialector {
	// 也可以使用内存模式
	// file::memory:?cache=shared
	// 开启WAL模式
	// file::./database.db?cache=shared&mode=rwc&_journal_mode=WAL
	// 参考 https://stackoverflow.com/questions/57118674/go-sqlite3-with-journal-mode-wal-gives-database-is-locked-error
	// dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc&_journal_mode=WAL", filename)
	dsn := filename

	// https://gorm.io/docs/connecting_to_the_database.html#SQLite
	// gorm.io/driver/sqlite 动态编译需要CGO
	// github.com/glebarez/sqlite 静态编译
	conn := sqlite.Open(dsn)
	return conn
}

//func newGormLogger(conf *viper.Viper, l *zap.Logger) *zapgorm2.Logger {
//	loggerLevel := gormlogger.Warn
//	if conf.GetBool("app.debug") {
//		loggerLevel = gormlogger.Info
//	}
//	logger := &zapgorm2.Logger{
//		ZapLogger:                 l.WithOptions(zap.WithCaller(false)),
//		LogLevel:                  loggerLevel,
//		SlowThreshold:             500 * time.Millisecond,
//		SkipCallerLookup:          true,
//		IgnoreRecordNotFoundError: true,
//		Context:                   nil,
//	}
//	logger.SetAsDefault()
//	return logger
//}
