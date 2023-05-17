package forms

type Question struct {
	Text    string `json:"text" xml:"text" form:"text" binding:"required"`
	Enabled bool   `json:"enabled" xml:"enabled" form:"enabled" binding:"required"`
	ImageID int64  `json:"image_id" xml:"image_id" form:"image_id"`
	TopicID int64  `json:"topic_id" xml:"topic_id" form:"topic_id" binding:"required"`
}
