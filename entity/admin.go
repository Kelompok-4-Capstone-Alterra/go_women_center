package entity

type Admin struct {
	Id       string `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string
	Username string `gorm:"unique"`
}
