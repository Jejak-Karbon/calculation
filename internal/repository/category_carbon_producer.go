package repository

import (
	"strings"
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type CategoryCarbonProducer interface {
	Find(ctx context.Context,payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.CategoryCarbonProducer, *dto.PaginationInfo, error)
}

type category_carbon_producer struct {
	Db *gorm.DB
}

func NewCategoryCarbonProducer(db *gorm.DB) *category_carbon_producer {
	return &category_carbon_producer{
		db,
	}
}

func (c *category_carbon_producer) Find(ctx context.Context,payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.CategoryCarbonProducer, *dto.PaginationInfo, error) {
	var category_carbon_producers []model.CategoryCarbonProducer
	var count int64

	query := c.Db.WithContext(ctx).Model(&model.CategoryCarbonProducer{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&category_carbon_producers).Error

	return category_carbon_producers, dto.CheckInfoPagination(paginate, count), err
}
