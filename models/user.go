package models

// User ...
type User struct {
	ID       uint64 `json:"id" form:"id" gorm:"unique" gorm:"primaryKey"`
	Login    string `json:"login" form:"login" gorm:"unique"`
	Password string `json:"password" form:"password"`
}
