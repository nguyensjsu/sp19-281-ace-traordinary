package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type user struct {
	gorm.Model
	userid      string
	password    string
	firstname   string
	lastname    string
	phonenumber string
}

type registration struct {
	gorm.Model
	userid           string
	password         string
	firstname        string
	lastname         string
	mailvalid        bool
	phonenumber      string
	verificationcode string
	timestamp        string
}

func DBConnection() {
	db, err = gorm.Open("sqlite", "test.db")
	if err != nil {
		log.Panic(err)
		log.Panic("Error while opening connection to Sqlite Database")
	}
	defer db.Close()
	db.AutoMigrate(&user{})
	db.AutoMigrate(&registration{})
}
