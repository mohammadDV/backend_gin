package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohammadDV/backend_gin/serilizers"
	"github.com/mohammadDV/backend_gin/services"
)

type AuthApi interface {
	Register(ctx *gin.Context)
}

type authApi struct {
	authService services.AuthService
}

func NewAuthApi(authService services.AuthService) AuthApi {
	return &authApi{
		authService: authService,
	}
}

func (a authApi) Register(ctx *gin.Context) {
	var registerRequest serilizers.RegisterRequest
	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := a.authService.AddUserService(registerRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
	return 
}
