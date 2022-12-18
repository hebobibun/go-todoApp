package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

func ReadConfig() *Config {
	err := godotenv.Load("local.env")
	if err != nil {
		fmt.Println("Error while reading .env", err.Error())
		return nil
	}

	var res Config
	res.DBUser = os.Getenv("DBUSER")
	res.DBPass = os.Getenv("DBPASS")
	res.DBHost = os.Getenv("DBHOST")
	res.DBName = os.Getenv("DBNAME")
	readData := os.Getenv("DBPORT")
	res.DBPort, err = strconv.Atoi(readData)
	if err != nil {
		fmt.Println("Error while converting DBPort", err.Error())
		return nil
	}

	return &res
}

func ConnectSQL(c Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error is happening : ", err.Error())
	}
	return db
}