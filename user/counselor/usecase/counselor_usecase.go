package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	Counselor "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/repository"
	Review "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/review/repository"
)

type CounselorUsecase interface {
	GetAll(offset, limit int) ([]counselor.GetAllResponse, error)
	GetTotalPages(limit int) (int, error)
	// GetById(id string) (counselor.GetByIdResponse, error)
	// CreateReview(input counselor.CreateReviewRequest) error
}

type counselorUsecase struct {
	counselorRepo Counselor.CounselorRepository
	reviewRepo Review.ReviewRepository
}

func NewCounselorUsecase(CounselorRepo Counselor.CounselorRepository, ReviewRepo Review.ReviewRepository ) CounselorUsecase {
	return &counselorUsecase{counselorRepo: CounselorRepo, reviewRepo: ReviewRepo}
}

func(u *counselorUsecase) GetAll(offset, limit int) ([]counselor.GetAllResponse, error) {

	counselors, err := u.counselorRepo.GetAll(offset, limit)

	if err != nil {
		return nil, err
	}

	return counselors, nil
}

func(u *counselorUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.counselorRepo.Count()

	if err != nil {
		return 0, err
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

// func(u *counselorUsecase) CreateReview(inputReview counselor.CreateReviewRequest) error {
	
// 	_, err := u.counselorRepo.GetById(inputReview.CounselorID)

// 	if err != nil {
// 		return counselor.ErrCounselorNotFound
// 	}

	
// 	newReview := entity.Review{
// 		CounselorID: inputReview.CounselorID,
// 		UserID: inputReview.UserID,
// 		Rating: inputReview.Rating,
// 		Comment: inputReview.Comment,
// 	}
	
// 	// Check if user already give review
	
// 	oldReview, err := u.reviewRepo.GetByUserIdAndCounselorId(inputReview.UserID, inputReview.CounselorID)

// 	if err == nil {
		
// 		// Update review
// 		newReview.ID = oldReview.ID
		
// 	}else {

// 		// Create new review
// 		uuid, _ := helper.NewGoogleUUID().GenerateUUID()
// 		newReview.ID = uuid
// 	}

// 	newReview.Rating = inputReview.Rating
// 	newReview.Comment = inputReview.Comment

// 	// fmt.Println("result new -> ",newReview)

// 	err = u.reviewRepo.Save(newReview)
	
// 	if err != nil {
// 		return counselor.ErrInternalServerError
// 	}

// 	// Update counselor rating
// 	// rating, err := u.reviewRepo.GetAverageRating(inputReview.CounselorID)

// 	// if err != nil {
// 	// 	return counselor.ErrInternalServerError
// 	// }

// 	// u.counselorRepo.Update(inputReview.CounselorID, entity.Counselor{
// 	// 	Rating: rating,
// 	// })

// 	return nil	
// }
