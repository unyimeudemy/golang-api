package api

import (
	"Ecom/cmd/services/product"
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

// NewAPIServer is a constructor function that creates a new instance of APIServer.
// It takes an address (addr) which is the address and port number to listen on,
// and a pointer to an already initialized sql.DB database connection.
// It returns a pointer to the newly created APIServer instance.
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}



// define a method in the server that will make the server to run
func (s *APIServer) Run() error{

	// Here we 
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//***************************  USER ****************************************
	// Here we initialize the repository which will be used to communicate 
	// the database
	userStore := user.NewStore(s.db)

	// Then we pass the repository into the controller as argument
	userHandler := user.NewHandler(userStore)

	// call the register method on the controller to register all the routes passing
	// through the subroute
	userHandler.RegisterRoutes(subrouter)

	//***************************  PRODUCTS ****************************************
	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	log.Println("server listening on port: ", s.addr)

	// Start the server and listen for requests on the specified address
	return http.ListenAndServe(s.addr, router)
}