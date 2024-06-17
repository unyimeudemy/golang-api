package product

import (
	"Ecom/types"
	"Ecom/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/products", h.HandleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.HandleCreateProducts).Methods(http.MethodPost)

}

func (h *Handler) HandleGetProducts(w http.ResponseWriter, r *http.Request){
	ps, err := h.store.GetProducts()

	if err != nil{
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) HandleCreateProducts(w http.ResponseWriter, r *http.Request){
	
	var payload types.CreateProductPayload;

	if err := utils.ParseJSON(r, &payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateProduct(types.CreateProductPayload{
		Name: payload.Name,
		Description: payload.Description,
		Image: payload.Image,
		Price: payload.Price,
		Quantity: payload.Quantity,
	})

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}