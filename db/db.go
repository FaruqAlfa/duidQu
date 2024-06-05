package db

import (
	"duidQu/config"
	"duidQu/model"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(){
	cfg := config.NewConfig()
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	//migrasi model ke database
	err = DB.AutoMigrate(&model.User{}, &model.Pemasukan{}, &model.Pengeluaran{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	log.Println("Successfully connected to the database and migrated")
}