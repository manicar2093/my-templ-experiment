package models

import (
	"time"

	uuid "github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"default:gen_random_uuid()" json:"id" mapstructure:"id" param:"id"`
	Email     string    `json:"email" mapstructure:"email" param:"email" form:"email" validate:"required"`
	Password  string    `json:"password" mapstructure:"password" param:"password" form:"password" validate:"required"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at" param:"created_at" form:"created_at"`
}
