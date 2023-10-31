package main

import (
	"time"

	"github.com/darielgaizta/rssagg/internal/database"
	"github.com/google/uuid"
)

// Custom serializer
// This module is made to customized JSON format from the generated

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"api_key"`
}

func serializeUser(dbUser database.User) User {
	// Take database.User as parameter and return the locally-defined User
	return User{
		ID:        uuid.UUID(dbUser.ID),
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}
