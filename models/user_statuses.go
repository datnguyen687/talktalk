package models

const (
	// Activated ...
	Activated int = 1
	// NotActivated ...
	NotActivated int = 2
)

// UserStatus ...
type UserStatus struct {
	ID          int    `json:"id" gorm:"column_name:id; type:int; not null; primaryKey"`
	Status      string `json:"status" gorm:"column_name:status; type:string; not null"`
	Description string `json:"description" gorm:"column_name:description; type:string"`
}

// TableName ...
func (UserStatus) TableName() string {
	return "user_statuses"
}
