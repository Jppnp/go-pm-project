package controllers

import (
	"encoding/json"
	"net/http"
	"project_manage/database"
	"project_manage/models"
)

func GetHealthcheck(w http.ResponseWriter, r *http.Request) {
	var healtcheck *models.Healthcheck
	database.DB.First(&healtcheck)
	response := map[string]string{"message": healtcheck.Name}
	json.NewEncoder(w).Encode(response)
}
