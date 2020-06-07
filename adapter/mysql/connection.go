package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wahyurudiyan/authentication/config"
)

// const dsnFormat = "%s:%s@tcp(%s:%s)/%s?parseTime=%s&loc=%s&charset=%s&autocommit=%s" //e.g. ==> user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=True&loc=Asia%2FJakarta&charset=utf8&autocommit=false
const dsnFormat = "%s:%s@tcp(%s:%s)/%s?parseTime=%s&charset=%s" //e.g. ==> user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=True&charset=utf8

func NewConnection(cfg *config.Config) (*sql.DB, error) {
	return connectionBegin(cfg)
}

func connectionBegin(cfg *config.Config) (*sql.DB, error) {
	dataSource := fmt.Sprintf(dsnFormat,
		cfg.MysqlUsername,
		cfg.MysqlPassword,
		cfg.MysqlHost,
		cfg.MysqlPort,
		cfg.MysqlDBName,
		cfg.MysqlParseTime,
		cfg.MysqlCharset,
	)

	maxIddleConns, _ := strconv.Atoi(cfg.MysqlPoolMaxIddle)
	maxOpenConns, _ := strconv.Atoi(cfg.MysqlPoolMaxOpen)
	maxConnLifetime, _ := time.ParseDuration(cfg.MysqlPoolMaxLifetime)

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIddleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(maxConnLifetime)

	return db, nil
}
