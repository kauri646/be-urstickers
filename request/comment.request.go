package request

type CommentCreateRequest struct {
	FilmId uint `json:"film_id" validate:"required"`
	Comment      string         `json:"comment" validate:"required"`
}