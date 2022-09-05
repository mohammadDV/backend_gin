package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mohammadDV/backend_gin/entity"
)

func Index(ctx *gin.Context) {
	name := ctx.Param("name")
	age := ctx.Query("age")
	address := ctx.PostForm("address")
	ctx.JSON(http.StatusOK, gin.H{
		"name":    name,
		"age":     age,
		"address": address,
	})
}

func Test(ctx *gin.Context) {
	var user entity.User

	// err := ctx.BindQuery(&user)
	validate := validator.New()
	ctx.BindQuery(&user)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// validator
	if err := validate.Struct(user); err != nil {
		for _, err_user := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err_user.Error()})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
	})
}
