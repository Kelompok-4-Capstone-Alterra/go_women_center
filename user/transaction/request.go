package transaction

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
)

type SendTransactionRequest struct {
	//TODO: check data type
	UserCredential      *helper.JwtCustomUserClaims
	CounselorID         string `json:"counselor_id"`
	CounselingTimeID    string
	CounselingTimeStart string
	CounselingMethod    string
	ValueVoucher        float64
	GrossPrice          float64
	TotalPrice          float64
}
