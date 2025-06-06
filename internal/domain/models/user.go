package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"default:gen_random_uuid()" json:"id" mapstructure:"id" param:"id"`
	Email     string    `json:"email" mapstructure:"email" param:"email" form:"email" validate:"required"`
	Password  string    `json:"password" mapstructure:"password" param:"password" form:"password" validate:"required|min_len:10"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at" mapstructure:"created_at" param:"created_at" form:"created_at"`
	Status    Status    `json:"status" mapstructure:"status" param:"status" form:"status" validate:"required|in:active,deactived,pending"`
}

type Status string

const (
	Active      Status = "active"
	Deactivated Status = "deactivated"
	Pending     Status = "pending"
)
