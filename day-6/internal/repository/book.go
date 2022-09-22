package repository

import (
	"agmc-day-6/internal/dto"
	"agmc-day-6/internal/model"
	"context"
	"strings"

	"gorm.io/gorm"
)

type Book interface {
	Create(ctx context.Context, data model.Book) error
	Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Book, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID int) (model.Book, error)
	Update(ctx context.Context, ID int, data map[string]interface{}) error
	Delete(ctx context.Context, ID int) error
}

type book struct {
	DB *gorm.DB
}

func NewBook(db *gorm.DB) *book {
	return &book{
		db,
	}
}

func (b *book) Create(ctx context.Context, data model.Book) error {
	return b.DB.WithContext(ctx).Model(&model.Book{}).Create(&data).Error
}

func (b *book) Find(ctx context.Context, payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.Book, *dto.PaginationInfo, error) {
	var books []model.Book
	var count int64

	query := b.DB.WithContext(ctx).Model(&model.Book{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ? ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&books).Error

	return books, dto.CheckInfoPagination(paginate, count), err
}

func (b *book) FindByID(ctx context.Context, ID int) (model.Book, error) {
	var book model.Book

	if err := b.DB.WithContext(ctx).Model(&model.Book{}).Where("id = ?", ID).First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (b *book) Update(ctx context.Context, ID int, data map[string]interface{}) error {

	err := b.DB.WithContext(ctx).Where("id = ?", ID).Model(&model.Book{}).Updates(data).Error
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

func (b *book) Delete(ctx context.Context, ID int) error {

	err := b.DB.WithContext(ctx).Where("id = ?", ID).Delete(&model.Book{}).Error
	return err
}
