package model

import (
	"time"

	"gorm.io/gorm"
)

func (CategoryCarbonProducer) TableName() string {
    return "categories_carbon_producer"
}

type CategoryCarbonProducer struct {
	ID       uint   	`gorm:"primarykey;autoIncrement"`
	Name     string 	`json:"name" gorm:"size:200;not null"`
	Image    string 	`json:"image" gorm:"size:200;not null"`
	Model
}

func (c *CategoryCarbonProducer) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	return
}

// BeforeUpdate is a method for struct User
// gorm call this method before they execute query
func (c *CategoryCarbonProducer) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	return
}