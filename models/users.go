package models

import (
	"time"
)

// User ...
type User struct {
	Email    string `json:"email" gorm:"column_name:email;type:string;not null; primaryKey"`
	Password string `json:"password" gorm:"column_name:password;type:string;not null"`

	CreatedAt      time.Time `json:"created_at" gorm:"column_name:created_at;not null"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"column_name:updated_at;not null"`
	UserStatusesID int       `json:"status_id" gorm:"column_name:status_id; not null"`

	// Relation
	Status UserStatus `gorm:"foreignKey:UserStatusesID"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}

// UserFilter ...
type UserFilter struct {
	Email    *string
	Preloads bool
}
