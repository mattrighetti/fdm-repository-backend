package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var Db *sql.DB

// DBConfig is a struct that contains necessary data to connect to a MYSql server instance
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// buildDBConfig returns a pointer to a DBConfig struct
func buildDBConfig() *DBConfig {
	hostname := getEnv("DB_HOST")
	username := getEnv("DB_USER")
	password := getEnv("DB_PASSWD")

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

func getEnv(envName string) string {
	value, exists := os.LookupEnv(envName)
	if !exists {
		log.Fatalf("%s not in env.", envName)
	}
	return value
}

// DbURL returns Database URL readable for GORM
func DbURL() string {
	dbConfig := buildDBConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
