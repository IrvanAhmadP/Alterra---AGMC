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

func UpdateUser(id string, user models.User) (int64, error) {
	result := config.DB.Model(&models.User{}).Where("id = ?", id).Updates(&user)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}
	return result.RowsAffected, nil
}

func DeleteUser(id string) (int64, error) {
	var users []models.User

	result := config.DB.Delete(&users, id)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}
	return result.RowsAffected, nil
}
