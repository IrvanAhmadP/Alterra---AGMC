package book

import (
	"agmc-day-6/internal/dto"
	"agmc-day-6/internal/factory"
	"agmc-day-6/internal/model"
	"agmc-day-6/internal/repository"
	"agmc-day-6/pkg/constant"
	res "agmc-day-6/pkg/util/response"
	"context"
)

type Service interface {
	Create(ctx context.Context, payload *dto.CreateBookRequest) (string, error)
	Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Book], error)
	FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Book, error)
	Update(ctx context.Context, ID int, payload *dto.UpdateBookRequest) (string, error)
	Delete(ctx context.Context, ID int) (string, error)
}

type service struct {
	BookRepository repository.Book
}

func NewService(f *factory.Factory) Service {
	return &service{f.BookRepository}
}

func (s *service) Create(ctx context.Context, payload *dto.CreateBookRequest) (string, error) {
	var book = model.Book{
		Title:  payload.Title,
		Author: payload.Author,
		Year:   payload.Year,
	}

	err := s.BookRepository.Create(ctx, book)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Find(ctx context.Context, payload *dto.SearchGetRequest) (*dto.SearchGetResponse[model.Book], error) {
	books, info, err := s.BookRepository.Find(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := new(dto.SearchGetResponse[model.Book])
	result.Datas = books
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *dto.ByIDRequest) (*model.Book, error) {
	data, err := s.BookRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &data, nil
}

func (s *service) Update(ctx context.Context, ID int, payload *dto.UpdateBookRequest) (string, error) {
	var data = make(map[string]interface{})

	if payload.Title != nil {
		data["title"] = payload.Title
	}
	if payload.Author != nil {
		data["author"] = payload.Author
	}
	if payload.Year != nil {
		data["year"] = payload.Year
	}

	err := s.BookRepository.Update(ctx, ID, data)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}

func (s *service) Delete(ctx context.Context, ID int) (string, error) {
	_, err := s.BookRepository.FindByID(ctx, ID)
	if err != nil {
		if err == constant.RecordNotFound {
			return "failed", res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return "failed", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	err = s.BookRepository.Delete(ctx, ID)
	if err != nil {
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return "success", nil
}
