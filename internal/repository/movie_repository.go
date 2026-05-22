package repository

import (
	"context"
	"database/sql"

	"github.com/eben-vranken/movie-log/internal/models"
)

type MovieRepository struct {
	db *sql.DB
}

func (mr MovieRepository) Create(ctx context.Context, movie models.Movie) (models.Movie, error) {
	err := mr.db.QueryRowContext(ctx, `INSERT INTO movies 
	(title,
	genre_id,
	director,
	release_date,
	runtime,
	rating) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`, movie.Title, movie.GenreID, movie.Director, movie.ReleaseDate, movie.Runtime, movie.Rating).Scan(&movie.ID)

	return movie, err
}

func (mr MovieRepository) GetAll(ctx context.Context) ([]models.Movie, error) {
	rows, err := mr.db.QueryContext(ctx, "SELECT id, title, genre_id, director, release_date, runtime, rating FROM movies")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.GenreID, &movie.Director, &movie.ReleaseDate, &movie.Runtime, &movie.Rating)

		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return movies, rows.Err()
}

func CreateNewMovieRepository(db *sql.DB) MovieRepository {
	t := new(MovieRepository)
	t.db = db
	return *t
}
