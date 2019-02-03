package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DataSource string `yaml:"datasource"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Database   string `yaml:"database"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
}

func Open() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}

	c := Config{
		os.Getenv("DATA_SOURCE"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DATABASE"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	}

	return gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User, c.Password,
		c.Host, c.Port, c.Database))
}
