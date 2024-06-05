package model

import "gorm.io/gorm"

type Pemasukan struct {
    gorm.Model
    DanaMasuk  int    `json:"dana_masuk"`
    Tanggal    string `json:"tanggal"`
    Keterangan string `json:"keterangan"`
    UserID     uint   `json:"user_id"`
}
