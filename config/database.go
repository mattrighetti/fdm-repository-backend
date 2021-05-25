package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var Db *sql.DB

// dbConfig is a struct that contains necessary data to connect to a MYSql server instance
type dbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// buildDBConfig returns a pointer to a dbConfig struct
func buildDBConfig() *dbConfig {
	hostname := getEnv("DB_HOST")
	username := getEnv("DB_USER")
	password := getEnv("DB_PASSWD")
	dbSchema := getEnv("DB_SCHEMA")

	fmt.Printf("Init DB config with variables\nHOST: %s\nUSER: %s\nPassword: %s\n",
		hostname,
		username,
		password,
	)

	dbConfig := dbConfig{
		Host:     hostname,
		Port:     3306,
		User:     username,
		Password: password,
		DBName:   dbSchema,
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
