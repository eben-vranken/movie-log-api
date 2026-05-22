package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eben-vranken/movie-log/internal/models"
	"github.com/eben-vranken/movie-log/internal/repository"
)

type MovieHandler struct {
	mr *repository.MovieRepository
}

func (mh *MovieHandler) Create(w http.ResponseWriter, req *http.Request) {
	var movie models.Movie

	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&movie)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("200 - Bad Request"))
		return
	}

	createdMovie, err := mh.mr.Create(req.Context(), movie)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createdMovie)

	if err != nil {
		log.Print(err)
		log.Print("500 - Internal server error")
	}
}

func (mh *MovieHandler) GetAll(w http.ResponseWriter, req *http.Request) {
	movies, err := mh.mr.GetAll(req.Context())

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(movies)

	if err != nil {
		log.Print(err)
		log.Print("500 - Internal server error")
	}
}

func CreateNewMovieHandler(mr repository.MovieRepository) MovieHandler {
	t := new(MovieHandler)
	t.mr = &mr
	return *t
}
