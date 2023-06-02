package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/review"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetAllByCounselorID(counselorId string, offset, limit int) ([]entity.Review, error)
	CountByCounselorId(counselorId string) (int, error)
	Save(review entity.Review) error
	GetByUserIdAndCounselorId(userId, counselorId string) (entity.Review, error)
	GetByCounselorId(counselorId string, offset, limit int) ([]review.GetByCounselorId, error)
}

type mysqlReviewRepository struct {
	DB *gorm.DB
}

func NewMysqlReviewRepository(db *gorm.DB) ReviewRepository {
	return &mysqlReviewRepository{DB: db}
}

func(r *mysqlReviewRepository) GetAllByCounselorID(counselorId string, offset, limit int) ([]entity.Review, error) {
	
	var reviews []entity.Review

	err := r.DB.Where("counselor_id = ?", counselorId).Offset(offset).Limit(limit).Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func(r *mysqlReviewRepository) CountByCounselorId(counselorId string) (int, error) {
	
	var totalData int64

	err := r.DB.Model(&entity.Review{}).Where("counselor_id = ?", counselorId).Count(&totalData).Error

	if err != nil {
		return 0, err
	}

	return int(totalData), nil
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

func(r *mysqlReviewRepository) GetByCounselorId(counselorId string, offset, limit int) ([]review.GetByCounselorId, error) {
	
	var reviews []review.GetByCounselorId

	err := r.DB.Model(&entity.Review{}).Where("counselor_id = ?", counselorId).Find(&reviews).Offset(offset).Limit(limit).Error

	if err != nil {
		return reviews, err
	}

	return reviews, nil
	
}
