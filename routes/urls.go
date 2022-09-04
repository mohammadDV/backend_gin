package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mohammadDV/backend_gin/api/v1"
	"github.com/mohammadDV/backend_gin/db"
	"github.com/mohammadDV/backend_gin/middleware"
	"gorm.io/gorm"
)

var postDB *gorm.DB = db.ConnectPostgres()

func Urls() *gin.Engine {

	defer db.ClosePostgres(postDB)

	r := gin.Default()
	r.Use(middleware.CORS())

	r.NoRoute(middleware.NotRouteHandler())
	r.HandleMethodNotAllowed = true
	r.NoMethod(middleware.NotMethodHandler())

	apiV1 := r.Group("api/v1")
	apiV1.GET("/:name", v1.Index)

	return r

	// apiV1 := r.Group("api/v1/"){
	// 	apiV1.GET("/:name",)
	// }
	// r.GET("/", v1.Index)
	// r.POST("/home/:name", v1.Index)
	// r.POST("/struct", v1.Test)
}
