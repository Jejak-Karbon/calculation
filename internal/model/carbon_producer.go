package model

import (
	"time"

	"gorm.io/gorm"
)

func (CarbonProducer) TableName() string {
    return "carbon_producer"
}

type CarbonProducer struct {
	ID       uint   	`gorm:"primarykey;autoIncrement"`
	Name     string 	`json:"name" gorm:"size:200;not null"`
	Image    string 	`json:"image" gorm:"size:200;not null"`
	CategoryCarbonProducerID       uint   	`json:"category_carbon_producer_id" gorm:"not null"`
	Model
}

func (c *CarbonProducer) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (c *CarbonProducer) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}