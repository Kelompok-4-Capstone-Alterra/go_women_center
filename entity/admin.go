package entity

type Admin struct {
	ID       string `gorm:"primaryKey;uniqueindex;not null"`
	Username string `gorm:"type:varchar(150);uniqueindex;not null"`
	Email    string `gorm:"type:varchar(150);uniqueindex;not null"`
	Password string `gorm:"type:varchar(64)"`
}
