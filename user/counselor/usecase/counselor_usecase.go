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
	GetAll(offset, limit int, topic string) ([]counselor.GetAllResponse, error)
	GetTotalPages(limit int) (int, error)
	GetById(id string) (counselor.GetByResponse, error)
	GetAllReview(id string, offset, limit int) ([]counselor.ReviewResponse, error)
	GetTotalPagesReview(id string, limit int) (int, error)
	CreateReview(input counselor.CreateReviewRequest) error
	GetTotalPagesSearch(search, topic string, limit int) (int, error)
	Search(search, topic string, offset, limit int) ([]counselor.GetAllResponse, error)
}

type counselorUsecase struct {
	counselorRepo Counselor.CounselorRepository
	reviewRepo Review.ReviewRepository
}

func NewCounselorUsecase(CounselorRepo Counselor.CounselorRepository,ReviewRepo Review.ReviewRepository) CounselorUsecase {
	return &counselorUsecase{counselorRepo: CounselorRepo, reviewRepo: ReviewRepo}
}

func(u *counselorUsecase) GetAll(offset, limit int, topic string) ([]counselor.GetAllResponse, error) {

	counselorsRes, err := u.counselorRepo.GetAll(offset, limit, topic)

	if err != nil {
		return nil, err
	}

	return counselorsRes, nil
}

func(u *counselorUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.counselorRepo.Count()

	if err != nil {
		return 0, err
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
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

func(u *counselorUsecase) GetTotalPagesReview(id string, limit int) (int, error) {
	
	totalData, err := u.reviewRepo.CountByCounselorId(id)

	if err != nil {
		return 0, counselor.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func(u *counselorUsecase) GetAllReview(id string, offset, limit int) ([]counselor.ReviewResponse, error) {

	// Check if counselor exist
	_, err := u.counselorRepo.GetById(id)

	if err != nil {
		return []counselor.ReviewResponse{}, counselor.ErrCounselorNotFound
	}
	
	reviews, err := u.reviewRepo.GetByCounselorId(id, offset, limit)

	if err != nil {
		return []counselor.ReviewResponse{}, counselor.ErrInternalServerError
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
		return []counselor.ReviewResponse{}, err
	}

	return reviewsRes, nil
}

func(u *counselorUsecase) GetTotalPagesSearch(search, topic string, limit int) (int, error) {
	
	totalData, err := u.counselorRepo.CountBySearch(search, topic)

	if err != nil {
		return 0, err
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func(u *counselorUsecase) Search(search, topic string, offset, limit int) ([]counselor.GetAllResponse, error) {
	
	counselorsRes, err := u.counselorRepo.Search(search, topic, offset, limit)

	if err != nil {
		return nil, counselor.ErrInternalServerError
	}

	return counselorsRes, nil
}