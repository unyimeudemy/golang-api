package api

import (
	"Ecom/cmd/services/user"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// create a server object that is used to start your server
type APIServer struct {
	addr string
	db   *sql.DB
}

// define a method that will collect the necessary data and start a new instance of the server
// addr is port number
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}



// define a method in the server that will make the server to run
func (s *APIServer) Run() error{
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("server listening on port: ", s.addr)

	return http.ListenAndServe(s.addr, router)
}