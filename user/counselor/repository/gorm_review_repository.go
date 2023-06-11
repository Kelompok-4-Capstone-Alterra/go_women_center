package repository

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetByCounselorId(counselorId string, offset, limit int) ([]entity.Review, int64, error)
	Save(review entity.Review) error
	GetByUserIdAndCounselorId(userId, counselorId string) (entity.Review, error)
	GetByUserId(userId string) (entity.Review, error)
}

type mysqlReviewRepository struct {
	DB *gorm.DB
}

func NewMysqlReviewRepository(db *gorm.DB) ReviewRepository {
	return &mysqlReviewRepository{DB: db}
}

func(r *mysqlReviewRepository) GetByCounselorId(counselorId string, offset, limit int) ([]entity.Review, int64, error) {
	
	var reviews []entity.Review
	var totalData int64
	err := r.DB.
		Model(&entity.Review{}).
		Where("counselor_id = ?", counselorId).
		Count(&totalData).
		Offset(offset).
		Limit(limit).
		Find(&reviews).Error

	if err != nil {
		return nil, totalData ,err
	}

	return reviews, totalData ,nil
}

func(r *mysqlReviewRepository) Save(review entity.Review) error {

	err := r.DB.Transaction(func(tx *gorm.DB) error {

		var avgRating float32

		err := tx.Model(&entity.Review{}).Where("counselor_id = ?", review.CounselorID).Select("AVG(rating)").Scan(&avgRating).Error
		
		if err != nil {
			return err
		}
		
		err = tx.Model(&entity.Counselor{}).Where("id = ?", review.CounselorID).Update("rating", avgRating).Error

		if err != nil {
			return err
		}

		err = tx.Save(&review).Error

		if err != nil {
			return err
		}

		return nil
	})

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

func(r *mysqlReviewRepository) GetByUserId(userId string) (entity.Review, error){
	review := entity.Review{}

	err := r.DB.Where("user_id = ?", userId).First(&review).Error
	
	if err != nil {
		return review, err
	}

	return review, nil
}
