package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PortHTTP string `envconfig:"PORT_HTTP" default:"8088"`

	// Mysql DSN Configuration
	MysqlUsername  string `envconfig:"MYSQL_USER" default:"root"`
	MysqlPassword  string `envconfig:"MYSQL_PASS" default:""`
	MysqlHost      string `envconfig:"MYSQL_HOST" default:"localhost"`
	MysqlPort      string `envconfig:"MYSQL_PORT" default:"3306"`
	MysqlDBName    string `envconfig:"MYSQL_DBNAME" default:"account"`
	MysqlParseTime string `envconfig:"MYSQL_PARSETIME" default:"True"`
	MysqlCharset   string `envconfig:"MYSQL_CHARSET" default:"utf8"`
	// MysqlLocation   string `envconfig:"MYSQL_LOCATION" default:"Asia/Jakarta"`
	// MysqlAutocommit string `envconfig:"MYSQL_AUTOCOMMIT" default:"false"`

	// Mysql Pool Configuration
	MysqlPoolMaxIddle    string `envconfig:"MYSQL_MAX_IDDLE" default:"2"`
	MysqlPoolMaxOpen     string `envconfig:"MYSQL_MAX_IDDLE" default:"10"`
	MysqlPoolMaxLifetime string `envconfig:"MYSQL_MAX_IDDLE" default:"30m"`
}

func GetConfiguration() *Config {
	var config Config
	envconfig.MustProcess("web_app", &config)

	return &config
}
