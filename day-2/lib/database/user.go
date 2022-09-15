package database

import (
	"agmc/config"
	"agmc/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id string) (interface{}, error) {
	var user models.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(user models.User) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(id string, user models.User) error {
	if err := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	var users []models.User

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return err
	}
	return nil
}
