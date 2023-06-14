package transaction

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type SendTransactionRequest struct {
	//TODO: check data type
	UserCredential      *helper.JwtCustomUserClaims
	CounselorID         string `json:"counselor_id"`
	CounselorTopicKey   int    `json:"counselor_topic"`
	CounselingTimeID    string `json:"counselor_time_id"`
	CounselingTimeStart string `json:"counseling_time_start"`
	CounselingMethod    string `json:"counseling_method"`
	ValueVoucher        int64  `json:"value_voucher"`
	GrossPrice          int64  `json:"gross_price"`
	TotalPrice          int64  `json:"total_price"`
}
