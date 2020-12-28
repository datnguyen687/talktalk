package models

import "time"

const (
	// ActivationCodeLength ...
	ActivationCodeLength int = 6
)

// ActivationCode ...
type ActivationCode struct {
	ID        int       `json:"id" gorm:"column_name:id; type:serial; not nulll; primaryKey"`
	Code      string    `json:"code" gorm:"column_name:code; type:string; not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column_name:created_at; not null"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column_name:expired_at; not null"`
	UserEmail string    `json:"user_email" gorm:"column_name:user_email; not null"`

	User User `gorm:"foreignKey:UserEmail"`
}

// TableName ...
func (ActivationCode) TableName() string {
	return "activation_codes"
}
