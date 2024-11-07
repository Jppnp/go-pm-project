package routes

import (
	"project_manage/controllers"

	"github.com/gin-gonic/gin"
)

func GinRoutes(router *gin.Engine) {
	router.GET("/healthcheck", controllers.GetHealthcheck)
}
