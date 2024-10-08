package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/userAdityaa/bookmark-backend/models"
	"github.com/userAdityaa/bookmark-backend/utils"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ProfileImage string `json:"profile_image"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		hashedPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			http.Error(w, "Could not hash password", http.StatusInternalServerError)
		}
		user := models.User{
			Username:     req.Username,
			Email:        req.Email,
			Password:     hashedPassword,
			ProfileImage: req.ProfileImage,
		}
		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Could not create user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	}
}

func LoginUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		if !utils.CheckPasswordHash(req.Password, user.Password) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		token, err := utils.GenerateToken(user.ID, user.Email)
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Login successful",
			"token":   token,
		})
	}
}
