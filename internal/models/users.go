package models

import "time"

// User ...
type User struct {
	Email     string     `json:"email" gorm:"column_name:email;type:string;not null;primaryKey"`
	CreatedAt time.Time  `json:"created_at" gorm:"column_name:created_at;type:timestamp;not null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column_name:deleted_at;type:timestamp"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column_name:updated_at;type:timestamp"`

	UserStatusID int `json:"user_status_id" gorm:"column_name:user_status_id;not null"`

	UserStatus UserStatus `json:"user_status" gorm:"foreignKey:UserStatusID"`
}

// TableName ...
func (User) TableName() string {
	return "users"
}
