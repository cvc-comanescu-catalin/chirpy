package main

import (
	"encoding/json"
	"net/http"

	"github.com/cvc-comanescu-catalin/chirpy/models"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.CreateUser(req.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create user", err)
		return
	}

	responseUser := models.User{
		ID: user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email: user.Email,
	}

	respondWithJSON(w, http.StatusCreated, responseUser)
}

// type User struct {
// 	ID        uuid.UUID `json:"id"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	Email     string    `json:"email"`
// }

// func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		Email string `json:"email"`
// 	}
// 	type response struct {
// 		User
// 	}

// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
// 		return
// 	}

// 	user, err := cfg.db.CreateUser(r.Context(), params.Email)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, response{
// 		User: User{
// 			ID:        user.ID,
// 			CreatedAt: user.CreatedAt,
// 			UpdatedAt: user.UpdatedAt,
// 			Email:     user.Email,
// 		},
// 	})
// }
