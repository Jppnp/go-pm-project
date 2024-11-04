package routes

import (
	"net/http"
	"project_manage/controllers"
)

func InitializeRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetHealthcheck(w, r)
		}
	})
}
