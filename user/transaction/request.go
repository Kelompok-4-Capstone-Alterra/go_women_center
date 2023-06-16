package transaction

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type SendTransactionRequest struct {
	//TODO: check data type
	UserCredential        *helper.JwtCustomUserClaims
	CounselorID           string `json:"counselor_id" validate:"required,uuid"`
	CounselorTopicKey     int    `json:"counselor_topic_key" validate:"required"`
	ConsultationDateID    string `json:"consultation_date_id" validate:"required,uuid"`
	ConsultationTimeID    string `json:"consultation_time_id" validate:"required,uuid"`
	ConsultationTimeStart string `json:"consultation_time_start" validate:"required"`
	ConsultationMethod    string `json:"consultation_method" validate:"required"`
	ValueVoucher          int64  `json:"value_voucher" validate:"required"`
	GrossPrice            int64  `json:"gross_price" validate:"required"`
	TotalPrice            int64  `json:"total_price" validate:"required"`
}

type UserJoinHandlerRequest struct {
	UserId        string `json:"user_id" validate:"required,uuid"`
	TransactionId string `json:"transaction_id" validate:"required,uuid"`
}
