package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"

	"productivity-pal/internal/models"
)

type Svc struct {
	DB *sql.DB
}

func (s *Svc) UserHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello world from my new server!"
	_, err := w.Write([]byte(msg))
	if err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("UserHandler executed successfully")
}

// SignupHandler inserts user, preferences, and health details
func (s *Svc) SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		http.Error(w, "First name, last name, and email are required", http.StatusBadRequest)
		return
	}

	// Insert into Users table
	err = s.DB.QueryRow(`
		INSERT INTO Users (FirstName, LastName, Email)
		VALUES ($1, $2, $3)
		RETURNING UserID
	`, user.FirstName, user.LastName, user.Email).Scan(&user.ID)
	if err != nil {
		http.Error(w, "Failed to insert user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate UUID for preferences
	prefID := uuid.New().String()

	// Insert into Preferences table
	_, err = s.DB.Exec(`
	INSERT INTO Preferences (PreferenceID, UserID, HealthAdvice, FinanceAdvice, DailyTasks)
	VALUES ($1, $2, $3, $4, $5)
`, prefID, user.ID, user.Preferences.HealthAdvice, user.Preferences.FinanceAdvice, user.Preferences.DailyTasks)
	if err != nil {
		http.Error(w, "Failed to insert preferences: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert into HealthDetails table and return HealthID
	err = s.DB.QueryRow(`
	INSERT INTO HealthDetails (UserID, Age, Height, Weight, ActivityLevel, HeartRate, GoalWeight, GeneralGoal)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING HealthID
`, user.ID, user.HealthDetails.Age, user.HealthDetails.Height, user.HealthDetails.Weight,
		user.HealthDetails.ActivityLevel, user.HealthDetails.HeartRate,
		user.HealthDetails.GoalWeight, user.HealthDetails.GeneralGoal).Scan(&user.HealthDetails.HealthID)
	if err != nil {
		http.Error(w, "Failed to insert health details: "+err.Error(), http.StatusInternalServerError)
		return
	}

	msg := "Signup completed successfully!"
	w.Header().Set("Content-Type", "application/json")
	log.Println("User signed up successfully:", user.Email)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"message": "` + msg + `"}`))
}
