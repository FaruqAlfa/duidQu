package db

import (
    "duidQu/config"
    "duidQu/model"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Init() {
    cfg := config.NewConfig()
    var err error
    DB, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error opening database: %v\n", err)
    }

    // Migrasi model ke database
    err = DB.AutoMigrate(&model.User{}, &model.Pemasukan{}, &model.Pengeluaran{})
    if err != nil {
        log.Fatalf("Error migrating database: %v\n", err)
    }

    log.Println("Successfully connected to the database and migrated")
}
