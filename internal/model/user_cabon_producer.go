package model

import (
	"time"

	"gorm.io/gorm"
)

func (UserCarbonProducer) TableName() string {
    return "user_carbon_producer"
}

type UserCarbonProducer struct {
	ID       uint   	`gorm:"primarykey;autoIncrement"`
	UserID       uint   	`json:"user_id" gorm:"not null"`
	CarbonProducerID       uint   	`json:"carbon_producer_id" gorm:"not null"`
	Amount	float32 `json:"amount"  gorm:"not null"`
	Model
}

func (u *UserCarbonProducer) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (u *UserCarbonProducer) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}