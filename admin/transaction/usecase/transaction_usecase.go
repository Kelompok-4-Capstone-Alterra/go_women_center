package usecase

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction"
	trRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/transaction/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	// "golang.org/x/sync/errgroup"
)

type TransactionUsecase interface {
	GetAll(search, sortBy string, offset, limit int) (code int, totalPages int, data []entity.Transaction, err error)
	SendLink(req transaction.SendLinkRequest) (code int, err error)
	CancelTransaction(req transaction.CancelTransactionRequest) (int, error)
	GetAllForReport(tReq transaction.ReportRequest) ([]entity.Transaction, int, error)
	GenerateReport(transactionRecord []entity.Transaction) (string, int, error)
}

type transactionUsecase struct {
	repo        trRepo.MysqlTransactionRepository
	voucherRepo trRepo.MysqlVoucherRepository
	uuidGen     helper.UuidGenerator
}

func NewtransactionUsecase(
	trRepo trRepo.MysqlTransactionRepository,
	voucherRepo trRepo.MysqlVoucherRepository,
	uuidGen helper.UuidGenerator,
) TransactionUsecase {
	return &transactionUsecase{
		repo:        trRepo,
		voucherRepo: voucherRepo,
		uuidGen:     uuidGen,
	}
}

func (tu *transactionUsecase) GetAll(search, sortBy string, offset, limit int) (int, int, []entity.Transaction, error) {
	// TODO: pagination
	switch sortBy {
	case "newest":
		sortBy = "created_at DESC"
	case "oldest":
		sortBy = "created_at ASC"
	}

	data, totalData, err := tu.repo.GetAll(search, sortBy, offset, limit)
	if err != nil {
		return http.StatusInternalServerError,
			0,
			nil,
			err
	}

	return http.StatusOK,
		helper.GetTotalPages(int(totalData), limit),
		data,
		nil
}

func (tu *transactionUsecase) SendLink(req transaction.SendLinkRequest) (int, error) {
	_, err := tu.repo.UpdateById(req.TransactionId, req.Link, "waiting")
	if err != nil {
		if err.Error() == transaction.ErrEmptySlice.Error() {
			return http.StatusBadRequest, transaction.ErrUpdate
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (tu *transactionUsecase) CancelTransaction(req transaction.CancelTransactionRequest) (int, error) {
	// TODO: implement rollback
	transactionData, err := tu.repo.GetById(req.TransactionId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	_, err = tu.repo.UpdateById(req.TransactionId, "-", "canceled")
	if err != nil {
		if err.Error() == transaction.ErrEmptySlice.Error() {
			return http.StatusBadRequest, transaction.ErrUpdate
		}
		return http.StatusInternalServerError, err
	}

	voucherId, err := tu.uuidGen.GenerateUUID()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	expDate := time.Now().AddDate(0, 1, 0)
	voucher := entity.Voucher{
		ID:      voucherId,
		UserId:  transactionData.UserId,
		Value:   transactionData.GrossPrice,
		ExpDate: expDate,
	}

	_, err = tu.voucherRepo.CreateVoucher(voucher)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (tu *transactionUsecase) GetAllForReport(tReq transaction.ReportRequest) ([]entity.Transaction, int, error) {
	switch tReq.SortBy {
	case "newest":
		tReq.SortBy = "created_at DESC"
	case "oldest":
		tReq.SortBy = "created_at ASC"
	}
	data, dataCount, err := tu.repo.GetAllForReport(tReq)
	if err != nil {
		return nil, 0, err
	}

	return data, int(dataCount), nil
}

func (tu *transactionUsecase) GenerateReport(transactionRecord []entity.Transaction) (string, int, error) {
	workdir, _ := os.Getwd()
	fileLocation := filepath.Join(workdir, "admin", "transaction", "report.csv")
	file, err := os.Create(fileLocation)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(fileLocation)

	csvW := csv.NewWriter(file)

	row := []string{"date_id",
		"time_id",
		"id",
		"user_id",
		"counselor_id",
		"link",
		"time_start",
		"consultation_method",
		"status",
		"value_voucher",
		"gross_price",
		"total_price",
		"is_reviewed",
		"created_at",
	}
	csvW.Write(row)
	for _, tr := range transactionRecord {
		row = []string{tr.TimeId,
			tr.ID,
			tr.UserId,
			tr.CounselorID,
			tr.Link,
			tr.TimeStart,
			tr.ConsultationMethod,
			tr.Status,
			strconv.Itoa(int(tr.ValueVoucher)),
			strconv.Itoa(int(tr.GrossPrice)),
			strconv.Itoa(int(tr.TotalPrice)),
			strconv.FormatBool(tr.IsReviewed),
			tr.Created_at.String(),
		}
		err := csvW.Write(row)
		if err != nil {
			log.Fatalln("failure to write csv data")
		}
	}
	csvW.Flush()

	_, err = os.Open(fileLocation)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	file.Close()

	return fileLocation, http.StatusOK, nil
}
