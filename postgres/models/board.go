package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Board struct {
	gorm.Model
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name  	  string    `gorm:"size:255;not null;unique" json:"name"`
	Active 	  bool      `gorm:"boolean;default:false;not null;unique" json:"active"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
