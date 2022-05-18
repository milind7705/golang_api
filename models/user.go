package models

import (
	"errors"
	"golang_api/config"

	"gorm.io/gorm"
)

func GetAllUsers(user *[]User) (err error) {
	if err := config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {

	var existing_user User
	var address Address
	var company Company
	config.DB.Preload("Address").Preload("Company").First(&existing_user, user.ID)
	if existing_user.Address.City == user.Address.City {
		user.Address = existing_user.Address
	} else {
		err := config.DB.Where("city = ?", user.Address.City).First(&address).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.DB.Create(&user.Address)
			address = user.Address
		}
		user.Address = address
		user.AddressID = address.ID

	}
	if existing_user.Company.Name == user.Company.Name {
		user.Company = existing_user.Company
	} else {
		err := config.DB.Where("name = ?", user.Company.Name).First(&company).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			config.DB.Create(&user.Company)
			company = user.Company
		}
		user.Company = company
		user.CompanyID = company.ID
	}
	config.DB.Model(&existing_user).Updates(user)
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
