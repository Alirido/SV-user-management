package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
