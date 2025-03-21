package handlers

import (
	"database/sql"
	"log/slog"
	"net/http"
)

type APIServer struct {
	address string
	mux     *http.ServeMux
	db      *sql.DB
	logger  *slog.Logger
}

func NewAPIServer(address string, db *sql.DB, logger *slog.Logger) *APIServer {
	return &APIServer{
		address: address,
		mux:     Routes(),
		db:      db,
		logger:  logger,
	}
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", http.NotFoundHandler().ServeHTTP)

	return mux
}

func (s *APIServer) Run() {

}
