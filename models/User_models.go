package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`

	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
