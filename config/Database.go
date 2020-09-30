package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// DB gorm Database object
var DB *gorm.DB

// DBConfig is a struct that contains necessary data to connect to a MYSql server instance
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig returns a pointer to a DBConfig struct
func BuildDBConfig() *DBConfig {
	hostname, exist := os.LookupEnv("DB_HOST")
	if !exist {
		panic("DB_HOST not in env...Exiting.")
	}

	username, exist := os.LookupEnv("DB_USER")
	if !exist {
		panic("DB_USER not in env.")
	}

	password, exist := os.LookupEnv("DB_PASSWD")
	if !exist {
		panic("DB_PASSWD not in env.")
	}

	fmt.Printf("Init DB config with variables\nHOST: %s\nUSER: %s\nPassword: %s\n",
		hostname,
		username,
		password,
	)

	dbConfig := DBConfig{
		Host:     hostname,
		Port:     3306,
		User:     username,
		Password: password,
		DBName:   "fdm-repository",
	}

	return &dbConfig
}

// DbURL returns Database URL readable for GORM
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
