package mapper

import (
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/donor/repository/model"
)

func ModelToStructDonorRegister(datas []*model.DonorRegisterModel) []*domain.DonorRegisterDTO {
	result := make([]*domain.DonorRegisterDTO, 0)
	for _, regis := range datas {
		data := &domain.DonorRegisterDTO{
			Id:               regis.Id,
			Code:             regis.Code,
			UserId:           regis.UserId,
			DonorSchedulleId: regis.DonorSchedulleId,
			IsApprove:        regis.IsApprove,
			StatusApprove:    regis.StatusApprove,
			DonorProof:       regis.DonorProof,
			Status:           regis.Status,
			DonorSchedulle: domain.DonorSchedulleDTO{
				Id:            regis.DonorSchedulle.Id,
				UnitId:        regis.DonorSchedulle.UnitId,
				PlaceName:     regis.DonorSchedulle.PlaceName,
				Address:       regis.DonorSchedulle.Address,
				Date:          regis.DonorSchedulle.Date,
				TimeStart:     regis.DonorSchedulle.TimeStart,
				TimeEnd:       regis.DonorSchedulle.TimeEnd,
				TypeSchedulle: regis.DonorSchedulle.TypeSchedulle,
			},
			User: domain.User{
				Id:            regis.User.Id,
				Name:          regis.User.Name,
				Email:         regis.User.Email,
				Phone:         regis.User.Phone,
				JobId:         regis.User.JobId.String,
				UnitId:        regis.User.UnitId.String,
				PlaceOfBirth:  regis.User.PlaceOfBirth,
				DateOfBirth:   regis.User.DateOfBirth,
				Gender:        regis.User.Gender,
				SubDistrictId: regis.User.SubDistrictId,
				VillageId:     regis.User.VillageId,
				Address:       regis.User.Address,
				PostalCode:    regis.User.PostalCode,
				Role:          regis.User.Role,
				MemberCode:    regis.User.MemberCode.String,
			},
			Unit: domain.UnitDTO{
				Id:   regis.Unit.Id,
				Name: regis.Unit.Name,
			},
		}
		result = append(result, data)
	}
	return result
}
