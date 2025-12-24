package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Id			uint `gorm:"primaryKey"`
	Username 	string `json:"username" gorm:"unique"`
	Kata_sandi 	string `json:"kata_sandi"`
	Role    	string `json:"role"`
	Nama_lengkap     string `json:"nama_lengkap"`
	Statusdata 	string `json:"statusdata"`
}
