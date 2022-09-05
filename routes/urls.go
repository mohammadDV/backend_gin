package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mohammadDV/backend_gin/api/v1"
	"github.com/mohammadDV/backend_gin/db"
	"github.com/mohammadDV/backend_gin/middleware"
	"github.com/mohammadDV/backend_gin/repository"
	"github.com/mohammadDV/backend_gin/services"
	"gorm.io/gorm"
)

var (
	postDB         *gorm.DB                  = db.ConnectPostgres()
	authRepository repository.AuthRepository = repository.NewAuthRepository(postDB)
	authService    services.AuthService      = services.NewAuthRepository(authRepository)
	authApi        v1.AuthApi                = v1.NewAuthApi(authService)
)

func Urls() *gin.Engine {

	// defer db.ClosePostgres(postDB)

	r := gin.Default()
	r.Use(middleware.CORS())

	r.NoRoute(middleware.NotRouteHandler())
	r.HandleMethodNotAllowed = true
	r.NoMethod(middleware.NotMethodHandler())

	superGroup := r.Group("/api/v1")
	{
		authGroup := superGroup.Group("/auth")
		{
			authGroup.POST("/register", authApi.Register)
			// authGroup.GET("/login", v1.Index)
		}

		// todoGroup := superGroup.Group("/todo")
		// {
		// 	todoGroup.GET("/register", v1.Index)
		// }
	}

	return r
}
