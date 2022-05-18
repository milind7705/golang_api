package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID   int    `gorm:"primaryKey;column:id;"`
	Name string `gorm:"column:name" json:"name" binding:"required"`
}

type Address struct {
	gorm.Model
	ID     int    `gorm:"primaryKey;column:id;"`
	Street string `gorm:"column:street" json:"street" binding:"required"`
	City   string `gorm:"column:city" json:"city" binding:"required"`
}

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;column:id;"`
	Name      string `gorm:"column:name" json:"name" binding:"required"`
	Email     string `gorm:"column:email" json:"email" binding:"required"`
	Phone     string `gorm:"column:phone" json:"phone" binding:"required"`
	AddressID int
	Address   Address
	CompanyID int
	Company   Company
}

// If we need custom table names, we can use this
// func (b *User) TableName() string {
// 	return "user"
// }

// func (b *Company) TableName() string {
// 	return "company"
// }

// func (b *Address) TableName() string {
// 	return "address"
// }
