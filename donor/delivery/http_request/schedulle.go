package http_request

type SchedulleStore struct {
	UnitId    int64  `json:"unitId" validate:"required"`
	PlaceName string `json:"placeName" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Date      string `json:"date" validate:"required"`
	TimeStart string `json:"timeStart" validate:"required"`
	TimeEnd   string `json:"timeEnd" validate:"required"`
	Type      string `json:"type" validate:"required"`
}
