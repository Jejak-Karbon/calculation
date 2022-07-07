package repository

import (
	"strings"
	"context"

	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type UserCarbonProducer interface {
	Find(ctx context.Context, payload *dto.SearchGetRequest, id uint, paginate *dto.Pagination) ([]model.UserCarbonProducer, *dto.PaginationInfo, error)
	Create(ctx context.Context, data model.UserCarbonProducer) error
}

type user_carbon_producer struct {
	Db *gorm.DB
}

func NewUserCarbonProducer(db *gorm.DB) *user_carbon_producer {
	return &user_carbon_producer{
		db,
	}
}

func (u *user_carbon_producer) Find(ctx context.Context, payload *dto.SearchGetRequest, id uint, paginate *dto.Pagination) ([]model.UserCarbonProducer, *dto.PaginationInfo, error) {
	var user_carbon_producer []model.UserCarbonProducer
	var count int64

	query := u.Db.WithContext(ctx).Model(&model.UserCarbonProducer{}).Where("user_id = ?  ", id)

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(name) LIKE ?  ", search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := dto.GetLimitOffset(paginate)

	err := query.Limit(limit).Offset(offset).Find(&user_carbon_producer).Error

	return user_carbon_producer, dto.CheckInfoPagination(paginate, count), err
}

func (u *user_carbon_producer) Create(ctx context.Context, data model.UserCarbonProducer) error {
	return u.Db.WithContext(ctx).Model(&model.UserCarbonProducer{}).Create(&data).Error
}