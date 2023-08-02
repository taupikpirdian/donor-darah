package model

type VillageModel struct {
	Id            string `json:"id"`
	Code          string `json:"code"`
	SubDistrictId string `json:"sub_district_id"`
	Name          string `json:"name"`
}
