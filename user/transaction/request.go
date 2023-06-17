package transaction

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type SendTransactionRequest struct {
	UserCredential        *helper.JwtCustomUserClaims
	CounselorID           string `json:"counselor_id" validate:"required,uuid"`
	ConsultationDateID    string `json:"consultation_date_id" validate:"required,uuid"`
	ConsultationTimeID    string `json:"consultation_time_id" validate:"required,uuid"`
	ConsultationTimeStart string `json:"consultation_time_start" validate:"required"`
	ConsultationMethod    string `json:"consultation_method" validate:"required"`
	VoucherId             string `json:"voucher_id"`
}

type UserJoinHandlerRequest struct {
	UserId        string `json:"user_id" validate:"required,uuid"`
	TransactionId string `json:"transaction_id" validate:"required,uuid"`
}
