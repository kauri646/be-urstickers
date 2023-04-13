package entity



type Film struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Judul      string `json:"judul" gorm:"foreignKey"`
	JenisFilm     string `json:"jenis_film"`
	Produser  string `json:"produser"`
	Sutradara   string `json:"sutradara"`
	Penulis     string `json:"penulis"`
	Produksi	string `json:"produksi"`
	Casts	string `json:"casts"`
	Sinopsis	string `json:"sinopsis"`
	Like 	uint `json:"like"`

	
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time	`json:"updated_at"`
	// DeleteAt gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}

type TheaterId struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	TheaterId uint `json:"theater_id"`
	Judul      string `json:"judul" gorm:"foreignKey"`
	JenisFilm     string `json:"jenis_film"`
	Produser  string `json:"produser"`
	Sutradara   string `json:"sutradara"`
	Penulis     string `json:"penulis"`
	Produksi	string `json:"produksi"`
	Casts	string `json:"casts"`
	Sinopsis	string `json:"sinopsis"`
	Like 	uint `json:"like"`
}




