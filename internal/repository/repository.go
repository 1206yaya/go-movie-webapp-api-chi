package repository

import "github.com/1206yaya/go-movie-webapp-api-chi/internal/models"

type DatabaseRepo interface {
	AllMovies() ([]*models.Movie, error)
}
