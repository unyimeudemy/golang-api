package user

import (
	"Ecom/cmd/services/auth"
	"Ecom/config"
	"Ecom/types"
	"Ecom/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

//********************** THIS IS THE CONTROLLER ****************************************

type Handler struct {
	store types.UserStore
}


// Here we are injecting the repository as a dependency into the controller 
// object called Handler. The dependency implements the UserStore interface
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// Registers all the routes with their corresponding request type
func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func(h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){
	//define a variable that will hold the incoming request body
	var payload types.LoginUserPayload

	// Here the code `utils.ParseJSON(r, payload)` populates the payload variable
	// with the request body from r.
	if err := utils.ParseJSON(r, &payload); err != nil {

		// And if an error occurs, a new response is created with w and is given
		// a status code of 400 (bad request) and message that is in err
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}


	//check if the user exist
	// NOTE store is a dependency we injected but you have to call it has
	// h.store.method
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"not found, invalid email or password",
		))
		return
	}


	if !auth.ComparePassword(u.Password, []byte(payload.Password)){
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"not found, invalid email or password",
		))
		return
	}

	secret := []byte(config.Envs.JWTSecret)

	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}


func(h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){

	//define a variable that will hold the incoming request body
	var payload types.RegisterUserPayload

	// Here the code `utils.ParseJSON(r, payload)` populates the payload variable
	// with the request body from r.
	if err := utils.ParseJSON(r, &payload); err != nil {

		// And if an error occurs, a new response is created with w and is given
		// a status code of 400 (bad request) and message that is in err
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}


	//check if the user exist
	// NOTE store is a dependency we injected but you have to call it has
	// h.store.method
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"user with email %s already exist",
			payload.Email,
		))
		return
	}

	hashedPassword, _ := auth.HashPassword(payload.Password)
	// hashedPassword, err := auth.HashPassword(payload.Password)

	//if user does not exist
	err = h.store.CreateUser(types.User{
		FirstName: payload.Firstname,
		LastName: payload.Lastname,
		Email: payload.Email,
		Password: hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}