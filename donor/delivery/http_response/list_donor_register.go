package http_response

import (
	"github.com/bxcodec/go-clean-arch/domain"
)

type ResponseDonorRegister struct {
	Id                              int64                              `json:"id"`
	Code                            string                             `json:"code"`
	Status                          string                             `json:"status"`
	BuktiDonor                      string                             `json:"bukti_donor"`
	ResponseDonorRegisterQuestioner []*ResponseDonorRegisterQuestioner `json:"donor_register_questioners"`
}

type ResponseDonorRegisterQuestioner struct {
	Id     int64  `json:"id"`
	Code   string `json:"code"`
	Answer string `json:"answer"`
}

type CustomReponseDonorRegisterList struct {
	Status *Status
	Data   []*ResponseDonorRegister
}

func MapResponseListDonorRegister(code int, message string, datas []*domain.DonorRegisterDTO) (*CustomReponseDonorRegisterList, error) {
	res := []*ResponseDonorRegister{}
	for _, data := range datas {
		// loop answer
		resAnswer := []*ResponseDonorRegisterQuestioner{}
		for _, dataAnswer := range data.DonorRegisterQuestionerDTO {
			addAnswer := &ResponseDonorRegisterQuestioner{
				Id:     dataAnswer.Id,
				Code:   dataAnswer.CodeQuestion,
				Answer: dataAnswer.Answer,
			}
			resAnswer = append(resAnswer, addAnswer)
		}
		add := &ResponseDonorRegister{
			Id:                              data.Id,
			Code:                            data.Code,
			Status:                          data.Status,
			BuktiDonor:                      data.DonorProof,
			ResponseDonorRegisterQuestioner: resAnswer,
		}
		res = append(res, add)
	}

	httpResponse := &CustomReponseDonorRegisterList{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: res,
	}

	return httpResponse, nil
}
