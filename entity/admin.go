package entity

type Admin struct {
	Id       string `gorm:"primaryKey"`
	Email    string
	Password string
	Username string
}