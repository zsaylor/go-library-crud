// Package config defines the database connection configurations used in package models.
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

//func Connect() {
//	d, err := gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/go-library?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		panic(err)
//	}
//	db = d
//}

//func InitDB(dataSource string) error {
//	var err error
//	db, err = gorm.Open("mysql", dataSource)
//	if err != nil {
//		return err
//	}
//	return err
//}

//func GetDB() *gorm.DB {
//	return db
//}
