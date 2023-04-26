package request

type FilmCreateRequest struct {
	Judul     string `json:"judul" validate:"required"`
	Thumbnail string `json:"thumbnail" validate:"required"`
	JenisFilm string `json:"jenisfilm" validate:"required"`
	Produser  string `json:"produser" validate:"required"`
	Sutradara string `json:"sutradara" validate:"required"`
	Penulis   string `json:"penulis" validate:"required"`
	Produksi  string `json:"produksi" validate:"required"`
	Casts     string `json:"casts" validate:"required"`
	Sinopsis  string `json:"sinopsis" validate:"required"`
	Like      uint   `json:"like"`
	
}

type FilmUpdateRequest struct {
	Judul     string `json:"judul" validate:"required"`
	Thumbnail string `json:"thumbnail" validate:"required"`
	JenisFilm string `json:"jenisfilm" validate:"required"`
	Produser  string `json:"produser" validate:"required"`
	Sutradara string `json:"sutradara" validate:"required"`
	Penulis   string `json:"penulis" validate:"required"`
	Produksi  string `json:"produksi" validate:"required"`
	Casts     string `json:"casts" validate:"required"`
	Sinopsis  string `json:"sinopsis" validate:"required"`
	Like      uint   `json:"like"`
	
}


