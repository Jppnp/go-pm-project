package controllers

import (
	"net/http"
	"project_manage/database"
	"project_manage/models"

	"github.com/gin-gonic/gin"
)

func GetHealthcheck(c *gin.Context) {
	var healtcheck *models.Healthcheck
	database.DB.First(&healtcheck)
	response := map[string]string{"message": healtcheck.Name}
	c.JSON(http.StatusOK, response)
}
