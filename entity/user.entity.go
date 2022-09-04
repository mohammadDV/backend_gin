package entity

type User struct {
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
}

type UserJson struct {
	FirstName string `json:"first_name" binding:"required" validate:"required,max=12"`
	LastName  string `json:"last_name" binding:"required" validate:"email,min=5"`
}
