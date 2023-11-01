package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/darielgaizta/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// Serializing JSON for request.
	type parameters struct {
		Name string `json:"name"`
	}

	// Decode JSON to allow Go digest!
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// Notice that CreateUserParams is an Any struct.
	// The attributes defined from required parameters in sql/queries/users.sql
	// Generated to internal/database/users.sql.go!
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	// [Old] For learning purposes, an empty payload which is an empty struct is passed.
	// Serializing JSON for response.
	// Return a response of a serialized data
	respondWithJSON(w, 201, serializeUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// Example for an authentication-required handler
	respondWithJSON(w, 200, serializeUser(user))
}
