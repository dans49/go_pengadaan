package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Nama 	string `json:"nama"`
	Email 	string `json:"email"`
	Alamat  string `json:"alamat"`
}
