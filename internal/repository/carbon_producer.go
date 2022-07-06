package repository

import (
	"strings"
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type CarbonProducer interface {
	Find(ctx context.Context,filter *dto.FilterCarbonProducer,payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.CarbonProducer, *dto.PaginationInfo, error)
	FindByID(ctx context.Context, ID uint) (model.CarbonProducer, error)
}

type carbon_producer struct {
	Db *gorm.DB
}

func NewCarbonProducer(db *gorm.DB) *carbon_producer {
	return &carbon_producer{
		db,
	}
}

func (c *carbon_producer) Find(ctx context.Context,filter *dto.FilterCarbonProducer,payload *dto.SearchGetRequest, paginate *dto.Pagination) ([]model.CarbonProducer, *dto.PaginationInfo, error) {
	var carbon_producers []model.CarbonProducer
	var count int64

	query := c.Db.WithContext(ctx).Model(&model.CarbonProducer{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	if filter.CategoryCarbonProducerID != ""{
		query = query.Where("category_carbon_producer_id = ?  ", filter.CategoryCarbonProducerID)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&carbon_producers).Error

	return carbon_producers, dto.CheckInfoPagination(paginate, count), err
}

func (c *carbon_producer) FindByID(ctx context.Context, ID uint) (model.CarbonProducer, error) {

	var data model.CarbonProducer
	err := c.Db.WithContext(ctx).Model(&data).Where("id = ?", ID).First(&data).Error

	return data, err
}


