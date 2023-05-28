package usecase

import (
	"mime/multipart"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/counselor"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/domain"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)


type counselorUsecase struct {
	counselorRepo domain.CounselorRepository
	reviewRepo domain.ReviewRepository
}

func NewCounselorUsecase(CRepo domain.CounselorRepository, RRepo domain.ReviewRepository ) domain.CounselorUsecase {
	return &counselorUsecase{counselorRepo: CRepo, reviewRepo: RRepo}
}

func(u *counselorUsecase) GetAll(offset, limit int) ([]domain.Counselor, error) {
	
	counselors, err := u.counselorRepo.GetAll(offset, limit)

	if err != nil {
		return nil, counselor.ErrInternalServerError
	}

	return counselors, nil
}

func(u *counselorUsecase) GetTotalPages(limit int) (int, error) {

	totalData, err := u.counselorRepo.Count()
	if err != nil {
		return 0, counselor.ErrInternalServerError
	}

	totalPages := helper.GetTotalPages(totalData, limit)

	return totalPages, nil
}

func(u *counselorUsecase) GetById(id string) (domain.Counselor, error) {
	
	counselorData, err := u.counselorRepo.GetById(id)

	if err != nil {
		return counselorData, counselor.ErrReviewNotFound
	}

	return counselorData, nil
}

func(u *counselorUsecase) Create(inputDetail counselor.CreateRequest, inputProfilePicture *multipart.FileHeader) error{
	
	_, err := u.counselorRepo.GetByEmail(inputDetail.Email)

	if err == nil {
		return counselor.ErrCounselorConflict
	}

	path, err := helper.UploadImageToS3(inputProfilePicture)

	if err != nil {
		return counselor.ErrInternalServerError
	}

	uuid, _ := helper.NewGoogleUUID().GenerateUUID()
	
	newCounselor := domain.Counselor{
		ID: uuid,
		Name: inputDetail.Name,
		Email: inputDetail.Email,
		Username: inputDetail.Username,
		Topic: inputDetail.Topic,
		Description: inputDetail.Description,
		Tarif: inputDetail.Tarif,
		ProfilePicture: path,
	}

	err = u.counselorRepo.Create(newCounselor)

	if err != nil {
		return counselor.ErrInternalServerError
	}
	
	return nil
}

func(u *counselorUsecase) Update(inputDetail counselor.UpdateRequest, inputProfilePicture *multipart.FileHeader) error {

	counselorData, err := u.counselorRepo.GetById(inputDetail.ID)
	
	if err != nil {
		return counselor.ErrReviewNotFound
	}

	_, err = u.counselorRepo.GetByEmail(inputDetail.Email)
	
	if err == nil {	
		if counselorData.Email != inputDetail.Email {
			return counselor.ErrEmailConflict
		}
	}

	counselorUpdate := domain.Counselor{
		Name: inputDetail.Name,
		Email: inputDetail.Email,
		Username: inputDetail.Username,
		Topic: inputDetail.Topic,
		Description: inputDetail.Description,
		Tarif: inputDetail.Tarif,
	}

	if inputProfilePicture != nil {
		err := helper.DeleteImageFromS3(counselorData.ProfilePicture)

		if err != nil {
			return counselor.ErrInternalServerError
		}
	
		path, err := helper.UploadImageToS3(inputProfilePicture)
		
		if err != nil {
			return counselor.ErrInternalServerError
		}
		
		counselorUpdate.ProfilePicture = path

	}
	
	err = u.counselorRepo.Update(counselorData.ID, counselorUpdate)

	if err != nil {
		return counselor.ErrInternalServerError
	}
	
	return nil
}

func(u *counselorUsecase) Delete(id string) error {
	
	counselorData, err := u.counselorRepo.GetById(id)

	if err != nil {
		return counselor.ErrReviewNotFound
	}
	
	err = u.counselorRepo.Delete(counselorData.ID)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	return nil
}

func(u *counselorUsecase) CreateReview(inputReview counselor.CreateReviewRequest) error {
	
	_, err := u.counselorRepo.GetById(inputReview.CounselorID)

	if err != nil {
		return counselor.ErrCounselorNotFound
	}

	
	newReview := domain.Review{
		CounselorID: inputReview.CounselorID,
		UserID: inputReview.UserID,
		Rating: inputReview.Rating,
		Comment: inputReview.Comment,
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
	newReview.Comment = inputReview.Comment

	// fmt.Println("result new -> ",newReview)

	err = u.reviewRepo.Save(newReview)
	
	if err != nil {
		return counselor.ErrInternalServerError
	}

	// Update counselor rating
	rating, err := u.reviewRepo.GetAverageRating(inputReview.CounselorID)

	if err != nil {
		return counselor.ErrInternalServerError
	}

	u.counselorRepo.Update(inputReview.CounselorID, domain.Counselor{
		Rating: rating,
	})

	return nil
}
