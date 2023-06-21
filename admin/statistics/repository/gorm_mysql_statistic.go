package repository

import (
	"gorm.io/gorm"
)

type StatisticRepo interface {
	GetRowCount(model interface{}) (int, error)
}

type statisticGormMysqlRepo struct {
	DB *gorm.DB
}

func NewStatisticGormMysqlRepo(db *gorm.DB) StatisticRepo {
	return &statisticGormMysqlRepo{
		DB: db,
	}
}

func (sr *statisticGormMysqlRepo) GetRowCount(model interface{}) (int, error) {
	count := int64(0)
	err := sr.DB.
		Model(&model).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return int(count), nil
}