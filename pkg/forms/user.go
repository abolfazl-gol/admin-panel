package forms

type Authenticate struct {
	Email    string `json:"email" xml:"email" form:"email" binding:"required,email"`
	Password string `json:"password" xml:"password" form:"password" binding:"required"`
}
