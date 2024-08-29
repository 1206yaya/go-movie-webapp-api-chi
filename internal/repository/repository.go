package repository

import (
	"database/sql"

	"github.com/1206yaya/go-movie-webapp-api-chi/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
