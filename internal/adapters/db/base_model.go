package db

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Conn interface {
	Create(value interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
}
