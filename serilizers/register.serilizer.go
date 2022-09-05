package serilizers

type RegisterRequest struct {
	FirstName string `json:"first_name" form:"firstname" binding:"required,min=1"`
	LastName  string `json:"last_name" form:"last_name" binding:"required,min=1"`
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string `json:"password" form:"password" binding:"required,min=4"`
}
