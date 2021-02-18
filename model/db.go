package model

import (
	"fmt"

	"github.com/2021-ZeroGravity-backend/log"

	"github.com/jinzhu/gorm"
	// MySQL driver.
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Fatal("Open database failed",
			zap.String("reason", err.Error()),
			zap.String("detail", fmt.Sprintf("Database connection failed. Database name: %s", name)))
	}

	return db
}

func (db *Database) Close() {
	DB.Self.Close()
}
