package models

import (
	"time"
)

type Profile struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Phone     string    `json:"phone" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(255)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Subscribe bool      `json:"subscribe" gorm:"type: boolean"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
