package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetAll(idCounselor string, offset, limit int) ([]entity.Review, error)
	Count(idCounselor string) (int, error)
	GetById(id string) (entity.Review, error)
	Save(review entity.Review) error
	GetAverageRating(idCounselor string) (float32, error)
	GetByUserIdAndCounselorId(userId, counselorId string) (entity.Review, error)
}

type mysqlReviewRepository struct {
	DB *gorm.DB
}

func NewMysqlReviewRepository(db *gorm.DB) ReviewRepository {
	return &mysqlReviewRepository{DB: db}
}

func(r *mysqlReviewRepository) GetAll(idCounselor string, offset, limit int) ([]entity.Review, error) {
	
	var reviews []entity.Review

	err := r.DB.Where("counselor_id = ?", idCounselor).Offset(offset).Limit(limit).Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func(r *mysqlReviewRepository) Count(idCounselor string) (int, error) {
	
	var totalData int64

	err := r.DB.Model(&entity.Review{}).Where("counselor_id = ?", idCounselor).Count(&totalData).Error

	if err != nil {
		return 0, err
	}

	return int(totalData), nil
}

func(r *mysqlReviewRepository) GetById(id string) (entity.Review, error) {
	
	var review entity.Review

	err := r.DB.Where("id = ?", id).First(&review).Error

	if err != nil {
		return review, err
	}

	return review, nil
}

func(r *mysqlReviewRepository) Save(review entity.Review) error {
	
	err := r.DB.Save(&review).Error

	if err != nil {
		return err
	}

	return nil
}

func(r *mysqlReviewRepository) GetByUserIdAndCounselorId(userId, counselorId string) (entity.Review, error){
	review := entity.Review{}

	err := r.DB.Where("user_id = ? AND counselor_id = ?", userId, counselorId).First(&review).Error
	
	if err != nil {
		return review, err
	}

	return review, nil
}

func(r *mysqlReviewRepository) GetAverageRating(idCounselor string) (float32, error) {
	
	var averageRating float32

	err := r.DB.Model(&entity.Review{}).Where("counselor_id = ?", idCounselor).Select("AVG(rating)").Scan(&averageRating).Error

	if err != nil {
		return 0, err
	}

	return averageRating, nil
}