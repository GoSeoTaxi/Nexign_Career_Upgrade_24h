package rest

import (
	"Nexign_Career_Upgrade_24h/internal/config"
	"Nexign_Career_Upgrade_24h/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func MainServer() {

	r := chi.NewRouter()
	r.Use(middleware.Compress(5, "gzip"))
	r.Post("/api/post/json", handlers.GetAPIJSON)
	http.ListenAndServe(":"+config.Port, r)

}
