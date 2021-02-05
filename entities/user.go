package entities

import (
	"time"
)

const (
	// UserBanned ...
	UserBanned int = -1
	// UserNotActivated ...
	UserNotActivated int = 0
	// UserActivated ...
	UserActivated int = 1
)

// UserDTO ...
type UserDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// User ...
type User struct {
	ID        int        `json:"id" gorm:"column_name:id;type:serial;autoIncrement;primaryKey"`
	Email     string     `json:"email" gorm:"column_name:email;type:VARCHAR(256);not null"`
	Password  string     `json:"password" gorm:"column_name:password;type:VARCHAR(256);not null"`
	Name      string     `json:"name" gorm:"column_name:name;type:VARCHAR(256)"`
	Status    int        `json:"status" gorm:"column_name:status;type:integer;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"column_name:created_at;not null;type:timestamp"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column_name:updated_at;type:timestamp"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}
