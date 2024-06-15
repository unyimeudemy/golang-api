package user

import (
	"Ecom/cmd/services/auth"
	"Ecom/types"
	"Ecom/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Here we injecting a dependency called store of type UserStore in the
// object Handler
type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func(h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){

}


func(h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){

	//get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//check if the user exist
	// NOTE store is a depency we injected but you have to call it has
	// h.store.method
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"user with email %s",
			payload.Email,
		))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

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