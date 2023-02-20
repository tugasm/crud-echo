package config

import (
	"crud-echo/models"
	"fmt"
	"log"
	"time"

	// "database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	if db != nil {
		return db
	}

	var err error
	dbConfig := Config.db
	// fmt.Printf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbConfig.UserDB, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)

	if dbConfig.Adapter == "postgres" {
		config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbConfig.Host, dbConfig.UserDB, dbConfig.Password, dbConfig.Name, dbConfig.Port)
		dsn := config
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		log.Println("Connected to Database Local postgressql")
	}

	if err != nil {
		log.Println("[Driver.ConnectDB] error when connect to database")
		log.Fatal("[Driver.ConnectDB] error when connect to database")
	} else {
		log.Println("SUCCESS CONNECT TO DATABASE")
	}
	sqldb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}


	pingDb := func (db *gorm.DB)  {
		sqldb.Ping()
	}

	go doEvery(6*time.Minute, pingDb, db)
	
	// Database Pooling
	sqldb.SetMaxIdleConns(20)
	sqldb.SetMaxOpenConns(200)
	sqldb.SetConnMaxLifetime(45 * time.Second)

  models.InitTableCustomer(db)

	return db
}

func doEvery(d time.Duration, f func(*gorm.DB), y *gorm.DB) {
	for _ = range time.Tick(d) {
		f(y)
	}
}
