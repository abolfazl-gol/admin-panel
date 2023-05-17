package forms

type Answer struct {
	Text       string `json:"text" xml:"text" form:"text" binding:"required"`
	Enabled    bool   `json:"enabled" xml:"enabled" form:"enabled" binding:"required"`
	ImageID    int64  `json:"image_id" xml:"image_id" form:"image_id"`
	Correct    bool   `json:"correct" xml:"correct" form:"correct"`
	QuestionID int64  `json:"question_id" xml:"question_id" form:"question_id" binding:"required"`
}
