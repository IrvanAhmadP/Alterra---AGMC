package repository

import (
	"agmc-day-6/internal/dto"
	"agmc-day-6/internal/model"
	"context"
	"strings"

	"gorm.io/gorm"
)

type User interface {
	Create(ctx context.Context, data model.User) error
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.User, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID int) (model.User, error)
	FindByEmail(ctx context.Context, email *string) (*model.User, error)
	Update(ctx context.Context, ID int, data map[string]interface{}) error
	Delete(ctx context.Context, ID int) error
}

type user struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (u *user) Create(ctx context.Context, data model.User) error {
	return u.DB.WithContext(ctx).Model(&model.User{}).Create(&data).Error
}

func (u *user) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.User, *dto.PaginationInfo, error) {
	var users []model.User
	var count int64

	query := u.DB.WithContext(ctx).Model(&model.User{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, dto.CheckInfoPagination(paginate, count), err
}

func (u *user) FindByID(ctx context.Context, ID int) (model.User, error) {
	var user model.User

	if err := u.DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", ID).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (u *user) FindByEmail(ctx context.Context, email *string) (*model.User, error) {
	var user model.User

	if err := u.DB.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) Update(ctx context.Context, ID int, data map[string]interface{}) error {

	err := u.DB.WithContext(ctx).Where("id = ?", ID).Model(&model.User{}).Updates(data).Error
	return err

	// var product model.ProductProduct

	// err := p.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	// if data.Name != nil {
	// product.Name = data.Name
	// }

	// product.Stock = data.Stock
	// product.Description = data.Description

	// err = p.Db.Save(&product).Error

	// return nil
}

func (u *user) Delete(ctx context.Context, ID int) error {

	err := u.DB.WithContext(ctx).Where("id = ?", ID).Delete(&model.User{}).Error
	return err
}
