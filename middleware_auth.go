package main

import (
	"fmt"
	"net/http"

	"github.com/darielgaizta/rssagg/internal/auth"
	"github.com/darielgaizta/rssagg/internal/database"
)

// Create a specific data type for the authentication-required handler
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("authentication error: %v", err))
			return
		}

		// NOTE THAT AT THIS POINT, THE API KEY SHOULD BE VALIDATED!
		// Notice that apiKey is required because in sql/queries/users.sql, GetUser takes api_key as argument.
		user, err := apiCfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("couldn't get user: %v", err))
			return
		}

		// Run the handler and pass the user as argument
		handler(w, r, user)
	}
}
