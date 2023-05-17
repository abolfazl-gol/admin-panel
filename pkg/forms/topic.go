package forms

type Topic struct {
	Name    string `json:"name" xml:"name" form:"name" binding:"required,min=2,max=100"`
	ImageID int64  `json:"image_id" xml:"image_id" form:"image_id"`
	Enabled bool   `json:"enabled" xml:"enabled" form:"enabled" binding:"omitempty"`
}
