package models

// UserStatus ...
type UserStatus struct {
	ID          int    `json:"id" gorm:"column_name:id;not null;type:integer;primaryKey"`
	Status      string `json:"status" gorm:"column_name:status;not null;type:string"`
	Description string `json:"description" gorm:"column_name:description;type:string"`
}

// TableName ...
func (UserStatus) TableName() string {
	return "user_statuses"
}
