package entity

type List struct{
	ID uint `json:"id" gorm:"primaryKey"`
	TheaterId uint `json:"theater_id"`
	Theater Theater `gorm:"foreignKey:TheaterId"`
	FilmId uint `json:"film_id"`
	Film Film `gorm:"foreignKey:FilmId"`
}