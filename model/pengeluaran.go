package model

import "gorm.io/gorm"

type Pengeluaran struct {
	gorm.Model
	DanaKeluar  int    `json:"dana_keluar"`
	Tanggal     string `json:"tanggal"`
	Keterangan  string `json:"keterangan"`
	UserID      uint   `json:"user_id"`
}