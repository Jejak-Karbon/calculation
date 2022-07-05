package repository

import (
	_ "strings"
	"context"

	_ "github.com/born2ngopi/alterra/basic-echo-mvc/internal/dto"
	"github.com/born2ngopi/alterra/basic-echo-mvc/internal/model"
	"gorm.io/gorm"
)

type UserCarbonProducer interface {
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

func (u *user_carbon_producer) Create(ctx context.Context, data model.UserCarbonProducer) error {
	return u.Db.WithContext(ctx).Model(&model.UserCarbonProducer{}).Create(&data).Error
}