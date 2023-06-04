package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	Counselor "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/repository"
	Review "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/review/repository"
	"golang.org/x/sync/errgroup"
)

type CounselorUsecase interface {
	GetAll(search, topic, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int, error)
	GetById(id string) (counselor.GetByResponse, error)
	GetAllReview(id string, offset, limit int) ([]counselor.ReviewResponse, int, error)
	CreateReview(input counselor.CreateReviewRequest) error
}

type counselorUsecase struct {
	counselorRepo Counselor.CounselorRepository
	reviewRepo Review.ReviewRepository
}

func NewCounselorUsecase(CounselorRepo Counselor.CounselorRepository,ReviewRepo Review.ReviewRepository) CounselorUsecase {
	return &counselorUsecase{counselorRepo: CounselorRepo, reviewRepo: ReviewRepo}
}

func(u *counselorUsecase) GetAll(search, topic, sortBy string, offset, limit int) ([]counselor.GetAllResponse, int, error) {

	switch sortBy {
	case "hight_price":
		sortBy = "price DESC"
	case "low_price":
		sortBy = "price ASC"
	default:
		sortBy = "rating DESC"
	}

	counselorsRes, totalData, err := u.counselorRepo.GetAll(search, topic, sortBy, offset, limit)
	
	if err != nil {
		return nil, 0, counselor.ErrInternalServerError
	}
	
	return counselorsRes, helper.GetTotalPages(int(totalData), limit), nil
}

func(u *counselorUsecase) GetById(id string) (counselor.GetByResponse, error) {
	
	counselorRes, err := u.counselorRepo.GetById(id)
	
	if err != nil {
		return counselorRes, counselor.ErrCounselorNotFound
	}

	return counselorRes, nil
}

func(u *counselorUsecase) CreateReview(inputReview counselor.CreateReviewRequest) error {
	
	_, err := u.counselorRepo.GetById(inputReview.CounselorID)

	if err != nil {
		return counselor.ErrCounselorNotFound
	}
	
	newReview := entity.Review{
		CounselorID: inputReview.CounselorID,
		UserID: inputReview.UserID,
		Rating: inputReview.Rating,
		Review: inputReview.Review,
	}
	
	// Check if user already give review
	
	oldReview, err := u.reviewRepo.GetByUserIdAndCounselorId(inputReview.UserID, inputReview.CounselorID)

	if err == nil {
		
		// Update review
		newReview.ID = oldReview.ID
		
	}else {

		// Create new review
		uuid, _ := helper.NewGoogleUUID().GenerateUUID()
		newReview.ID = uuid
	}

	newReview.Rating = inputReview.Rating
	newReview.Review = inputReview.Review

	err = u.reviewRepo.Save(newReview)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	return nil	
}

func(u *counselorUsecase) GetAllReview(id string, offset, limit int) ([]counselor.ReviewResponse, int, error) {

	// Check if counselor exist
	_, err := u.counselorRepo.GetById(id)

	if err != nil {
		return []counselor.ReviewResponse{}, 0, counselor.ErrCounselorNotFound
	}
	
	reviews, totalData, err := u.reviewRepo.GetByCounselorId(id, offset, limit)

	if err != nil {
		return []counselor.ReviewResponse{}, 0, counselor.ErrInternalServerError
	}

	var reviewsRes = make([]counselor.ReviewResponse, len(reviews))
	var g errgroup.Group

	for i, review := range reviews {
		i := i
		review := review
		g.Go(func () error{
			// user, err := u.reviewRepo.GetUserById(reviews.UserID)
			// if err != nil {
			// 	return err
			// }
			reviewRes := counselor.ReviewResponse{
				ID: review.ID,
				// ProfilePicture: user.ProfilePicture,
				// Name: user.Name,
				Rating: review.Rating,
				Review: review.Review,
				CreatedAt: review.CreatedAt.Format("2006-01-02 15:04:05"),
			}
			reviewsRes[i] = reviewRes

			return nil
		})
	}	

	if err := g.Wait(); err != nil {
		return []counselor.ReviewResponse{}, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), limit)

	return reviewsRes, totalPages ,nil
}