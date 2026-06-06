package order

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/havilz/ecommerce-backend/pkg/middleware"
	"github.com/havilz/ecommerce-backend/pkg/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// Checkout godoc
// @Summary      Checkout shopping cart
// @Description  Create an order from the user's current shopping cart and clear the cart
// @Tags         orders
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body CheckoutRequest true "Checkout Payload"
// @Success      201 {object} response.JSONResponse{data=OrderResponse}
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Router       /orders/checkout [post]
func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	order, err := h.service.Checkout(userID, req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, http.StatusCreated, "Order created successfully", order)
}

// GetOrders godoc
// @Summary      Get user orders
// @Description  Retrieve the order purchase history of the authenticated user
// @Tags         orders
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.JSONResponse{data=[]OrderResponse}
// @Failure      401 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /orders [get]
func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	orders, err := h.service.GetOrders(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch orders")
		return
	}
	response.Success(w, http.StatusOK, "Orders fetched successfully", orders)
}

// GetOrderByID godoc
// @Summary      Get order by ID
// @Description  Retrieve the details of a specific order by order ID
// @Tags         orders
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Order ID"
// @Success      200 {object} response.JSONResponse{data=OrderResponse}
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      404 {object} response.JSONResponse
// @Router       /orders/{id} [get]
func (h *Handler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	order, err := h.service.GetOrderByID(uint(id), userID)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "Order fetched successfully", order)
}
