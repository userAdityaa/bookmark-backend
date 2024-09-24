package routes

import (
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/userAdityaa/bookmark-backend/controllers"
	"gorm.io/gorm"
)

func SetUpRoutes(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)
	r.Post("/register", controllers.RegisterUser(db))
	r.Post("/login", controllers.LoginUser(db))
	return r
}
