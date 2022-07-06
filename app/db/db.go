package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB(DBDriver string) *gorm.DB {
	var err error
	var db *gorm.DB
	if DBDriver == "mysql" {
		DBUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DBUser"), os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBName"))
		DBConfig := mysql.Config{
			DSN:                       DBUrl, // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}

		db, err = gorm.Open(mysql.New(DBConfig), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DBDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DBDriver)
		}
	} else if DBDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DBHost"), os.Getenv("DBPort"), os.Getenv("DBUser"), os.Getenv("DBName"), os.Getenv("DBPassword"))
		db, err = gorm.Open(postgres.Open(DBURL))
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DBDriver)
			log.Fatal("This is the error connecting to postgres:", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", DBDriver)
		}
	} else if DBDriver == "sqlite" {
		db, err = gorm.Open(sqlite.Open("../../test.db"))
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DBDriver)
			log.Fatal("This is the error connecting to sqlite")
		}
	} else {
		fmt.Println("Unknown Driver")
		log.Fatal("This Driver is Unknown")
	}

	return db
}
