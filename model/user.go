package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name        string       `json:"name"`
    Username    string       `json:"username" gorm:"unique"`
    Email       string       `json:"email" gorm:"unique"`
    Password    string       `json:"password"`
    Pemasukan   []Pemasukan  `gorm:"foreignKey:UserID"`
    Pengeluaran []Pengeluaran `gorm:"foreignKey:UserID"`
}
