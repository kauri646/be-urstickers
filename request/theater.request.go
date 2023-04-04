package request

type TheaterCreateRequest struct {
	Kota	string `json:"kota" validate:"required"`
	Theater      string `json:"theater" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	Film	string `json:"film" validate:"required"`
	
}

type TheaterUpdateRequest struct {
	Kota	string `json:"kota"`
	Theater      string `json:"theater"`
	Phone     string `json:"phone"`
	Film	string `json:"film"`
	
}