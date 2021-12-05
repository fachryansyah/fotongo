package entities

import (
	"time"

	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "user"
}

type User struct {
	ID        string         `json:"id" gorm:"column:id;type:uuid;primarykey"`
	Name      string         `json:"name" gorm:"column:name"`
	Email     string         `json:"email" gorm:"column:email"`
	AvatarURL string         `json:"avatar_url" gorm:"column:avatar_url"`
	Role      string         `json:"role" gorm:"column:role"`
	Phone     string         `json:"phone" gorm:"column:phone"`
	Status    string         `json:"status" gorm:"column:status"`
	GoogleID  string         `json:"google_id" gorm:"column:google_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
