package main

// create a simple api route where the Api fetches all doctor information

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	// "github.com/gorilla/mux"
)

var db *sql.DB // Assume DB connection is already initialized

type Doctor struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Speciality string `json:"speciality"`
	Experience int    `json:"experience"`
}

// -----------------------------------
// Authentication Middleware
// -----------------------------------

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		// Check if Authorization header exists
		if authHeader == "" {
			http.Error(w, "Authorization token missing", http.StatusUnauthorized)
			return
		}

		// Optional: Check Bearer format
		tokenParts := strings.Split(authHeader, " ")

		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		// Token validation skipped intentionally
		// We are only checking existence/format

		next.ServeHTTP(w, r)
	})
}

// -----------------------------------
// Get Doctors API
// -----------------------------------

func GetDoctors(w http.ResponseWriter, r *http.Request) {

	// Default pagination values
	page := 1
	limit := 10

	// Read query params
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	if pageParam != "" {
		p, err := strconv.Atoi(pageParam)
		if err == nil && p > 0 {
			page = p
		}
	}

	if limitParam != "" {
		l, err := strconv.Atoi(limitParam)
		if err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	// SQL Query with Pagination
	query := `
		SELECT id, name, speciality, experience
		FROM doctors
		ORDER BY id
		LIMIT ? OFFSET ?
	`

	rows, err := db.Query(query, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var doctors []Doctor

	for rows.Next() {
		var d Doctor

		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Speciality,
			&d.Experience,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		doctors = append(doctors, d)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(doctors)
}

// -----------------------------------
// Main Function
// -----------------------------------

// func main() {
// 	r := mux.NewRouter()

// 	// Protected Route
// 	r.Handle("/doctors",
// 		AuthMiddleware(http.HandlerFunc(GetDoctors)),
// 	).Methods("GET")

// 	http.ListenAndServe(":8080", r)
// }
