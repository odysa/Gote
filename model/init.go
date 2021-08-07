package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *DataBase

type DataBase struct {
	Self *gorm.DB
}

func (db *DataBase) Init() {
	DB = &DataBase{Self: initDB()}
}

func (db *DataBase) Close() {
	db.Close()
}

func initDB() *gorm.DB {
	var db *gorm.DB

	if viper.GetString("db.type") == "mysql" {
		userName := viper.GetString("db.username")
		password := viper.GetString("db.password")
		address := viper.GetString("db.address")
		dbName := viper.GetString("db.dbname")
		db = openMysql(userName, password, address, dbName)
	}

	return db
}

func openMysql(username, password, address, dbName string) *gorm.DB {

	path := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		address,
		dbName,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})

	if err != nil {
		log.Fatalf("cannot open database %s", dbName)
	}

	return db
}
