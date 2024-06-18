package cart

import (
	"Ecom/types"
	"Ecom/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
	productStore types.ProductStore
	userStore types.UserStore
}

func NewHandler(
	store types.OrderStore, 
	productStore types.ProductStore, 
	userStore types.UserStore,
) *Handler{
	return &Handler{
		store: store,
		productStore: productStore,
		userStore: userStore,
	}
}

func(h *Handler) RegisterRoutes(router mux.Router){
	router.HandleFunc("/cart/checkout", h.HandleCheckout).Methods(http.MethodPost)
}

func(h *Handler) HandleCheckout(w http.ResponseWriter, r *http.Request){
	var cart types.CartCheckoutPayload
	userID := 0;   

	if err := utils.ParseJSON(r, &cart); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil{
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf(
			"invalid payload %v", errors,
		))
		return
	}


	//get product ids
	productIDs, err := getCardItemsIDs(cart.Items)
	if err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// fetch all products with id (returns an array of products)
	ps, err := h.productStore.GetProductsByIDs(productIDs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(ps, cart.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"total_price": totalPrice,
		"order_id":    orderID,
	})
}