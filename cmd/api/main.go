package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/eben-vranken/movie-log/internal/database"
	"github.com/eben-vranken/movie-log/internal/handlers"
	"github.com/eben-vranken/movie-log/internal/repository"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

var databaseURL = os.Getenv("DATABASE_URL")

func main() {
	http.HandleFunc("GET /health", loggingMiddleware(healthCheck))

	db, err := database.New(databaseURL)

	if err != nil {
		log.Panicf("Error when opening database %v\n", err)
	}

	movieRepository := repository.CreateNewMovieRepository(db)
	movieHandler := handlers.CreateNewMovieHandler(movieRepository)

	// Movie routes
	http.HandleFunc("POST /movies", loggingMiddleware(movieHandler.Create))

	log.Print("Listening to port 8080...")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func healthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Up and running!"))
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Println(req.URL.Path, "Initializing logging middleware")
		start := time.Now()

		recorder := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(recorder, req)
		duration := time.Since(start)
		log.Printf("[%s] %s %s %d", req.Method, req.RequestURI, duration, recorder.statusCode)
	})
}
