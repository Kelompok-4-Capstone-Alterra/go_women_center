package usecase

import (
	"log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	User "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	Counselor "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/repository"
	"golang.org/x/sync/errgroup"
)

type CounselorUsecase interface {
	GetAll(search, topic, sortBy string) ([]counselor.GetAllResponse, error)
	GetById(id string) (counselor.GetByResponse, error)
	GetAllReview(id string, offset, limit int) ([]counselor.GetAllReviewResponse, int, error)
	CreateReview(input counselor.CreateReviewRequest) error
}

type counselorUsecase struct {
	counselorRepo Counselor.CounselorRepository
	reviewRepo Counselor.ReviewRepository
	userRepo User.UserRepository
}

func NewCounselorUsecase(CounselorRepo Counselor.CounselorRepository, ReviewRepo Counselor.ReviewRepository, UserRepo User.UserRepository) CounselorUsecase {
	return &counselorUsecase{counselorRepo: CounselorRepo, reviewRepo: ReviewRepo, userRepo: UserRepo}
}

func(u *counselorUsecase) GetAll(search, topic, sortBy string) ([]counselor.GetAllResponse, error) {

	switch sortBy {
		case "highest_price":
			sortBy = "price DESC"
		case "lowest_price":
			sortBy = "price ASC"
		case "highest_rating":
			sortBy = "rating DESC"
	}

	counselorsRes, err := u.counselorRepo.GetAll(search, topic, sortBy)
	
	if err != nil {
		return nil, counselor.ErrInternalServerError
	}
	
	return counselorsRes, nil
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

func(u *counselorUsecase) GetAllReview(id string, offset, limit int) ([]counselor.GetAllReviewResponse, int, error) {

	// Check if counselor exist
	_, err := u.counselorRepo.GetById(id)

	if err != nil {
		return []counselor.GetAllReviewResponse{}, 0, counselor.ErrCounselorNotFound
	}
	
	reviews, totalData, err := u.reviewRepo.GetByCounselorId(id, offset, limit)

	if err != nil {
		log.Print(err.Error())
		return []counselor.GetAllReviewResponse{}, 0, counselor.ErrInternalServerError
	}

	var reviewsRes = make([]counselor.GetAllReviewResponse, len(reviews))
	var g errgroup.Group

	for i, review := range reviews {
		i := i
		review := review
		g.Go(func () error{
			user, err := u.userRepo.GetById(review.UserID)
			if err != nil {
				return counselor.ErrInternalServerError
			}
			reviewRes := counselor.GetAllReviewResponse{
				ID: review.ID,
				ProfilePicture: user.ProfilePicture,
				Username: user.Username,
				Rating: review.Rating,
				Review: review.Review,
				CreatedAt: review.CreatedAt.Format("2006-01-02 15:04:05"),
			}
			reviewsRes[i] = reviewRes

			return nil
		})
	}	

	if err := g.Wait(); err != nil {
		return []counselor.GetAllReviewResponse{}, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), limit)

	return reviewsRes, totalPages ,nil
}