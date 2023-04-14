package entity

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FilmId    uint      `json:"film_id"`
	Film Film `gorm:"-"` //foreignKey:FilmReference
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}