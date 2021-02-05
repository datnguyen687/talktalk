package entities

import "time"

// ActivationCode ...
type ActivationCode struct {
	Code      string    `json:"code" gorm:"column_name:code;type:VARCHAR(256);not null;primaryKey"`
	UserID    int       `json:"user_id" gorm:"column_name:user_id;type:integer;not null;primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"column_name:created_at;type:timestamp;not null"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column_name:expired_at;type:timestamp;not null"`

	User User `json:"user" gorm:"foreignKey:UserID"`
}

// TableName ...
func (ActivationCode) TableName() string {
	return "activation_codes"
}
