package entity

type Time struct {
	ID           string        `gorm:"primary_key;type:varchar(36);uniqueindex;not null"`
	CounselorID  string        `gorm:"type:varchar(36);not null"`
	Time         string        `gorm:"type:time(0);not null"`
	Transactions []Transaction `gorm:"foreignKey:UserId"`
}
