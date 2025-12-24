package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Nama_item 	string `json:"nama_item"`
	Stok 	string `json:"stok"`
	Harga  string `json:"harga"`
}
