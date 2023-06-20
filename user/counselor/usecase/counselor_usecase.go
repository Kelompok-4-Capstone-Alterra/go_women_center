package usecase

import (
	"log"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	User "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/auth/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor"
	Counselor "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/counselor/repository"
	Transaction "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/repository"
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
	transRepo Transaction.MysqlTransactionRepository
}

func NewCounselorUsecase(
		CounselorRepo Counselor.CounselorRepository,
		ReviewRepo Counselor.ReviewRepository,
		UserRepo User.UserRepository,
		TransactionRepo Transaction.MysqlTransactionRepository,
	) CounselorUsecase {
	return &counselorUsecase{CounselorRepo, ReviewRepo, UserRepo, TransactionRepo}
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
		if err.Error() == "record not found" {
			return counselor.ErrCounselorNotFound
		}
		return counselor.ErrInternalServerError
	}

	transaction, err := u.transRepo.GetById(inputReview.TransactionID)

	if err != nil {
		if err.Error() == "record not found" {
			return counselor.ErrTransactionNotFound
		}
		return counselor.ErrInternalServerError
	}

	if transaction.IsReviewed {
		return counselor.ErrReviewAlreadyExist
	}

	uuid, _ := helper.NewGoogleUUID().GenerateUUID()

	newReview := entity.Review{
		ID: uuid,
		CounselorID: inputReview.CounselorID,
		UserID: inputReview.UserID,
		Rating: inputReview.Rating,
		Review: inputReview.Review,
	}
	
	err = u.reviewRepo.Create(inputReview.TransactionID, newReview)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	return nil	
}

func(u *counselorUsecase) GetAllReview(id string, offset, limit int) ([]counselor.GetAllReviewResponse, int, error) {

	// Check if counselor exist
	_, err := u.counselorRepo.GetById(id)

	if err != nil {
		return nil, 0, counselor.ErrCounselorNotFound
	}
	
	reviews, totalData, err := u.reviewRepo.GetByCounselorId(id, offset, limit)

	if err != nil {
		log.Print(err.Error())
		return nil, 0, counselor.ErrInternalServerError
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
				UserProfile: user.ProfilePicture,
				Username: user.Username,
				Rating: review.Rating,
				Review: review.Review,
				CreatedAt: review.CreatedAt,
			}
			reviewsRes[i] = reviewRes

			return nil
		})
	}	

	if err := g.Wait(); err != nil {
		return nil, 0, err
	}

	totalPages := helper.GetTotalPages(int(totalData), limit)

	return reviewsRes, totalPages ,nil
}