package cart

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

// GetCart godoc
// @Summary      Get user cart
// @Description  Retrieve the current items and summary of the authenticated user's cart
// @Tags         cart
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.JSONResponse{data=CartResponse}
// @Failure      401 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /cart [get]
func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	cart, err := h.service.GetCart(userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch cart")
		return
	}
	response.Success(w, http.StatusOK, "Cart fetched successfully", cart)
}

// AddItem godoc
// @Summary      Add item to cart
// @Description  Add a specified quantity of a product to the user's shopping cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body AddCartRequest true "Add Cart Payload"
// @Success      201 {object} response.JSONResponse
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Router       /cart [post]
func (h *Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	var req AddCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := h.service.AddItem(userID, req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, http.StatusCreated, "Item added to cart", nil)
}

// UpdateItem godoc
// @Summary      Update cart item quantity
// @Description  Update the quantity of a specific item in the shopping cart
// @Tags         cart
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Cart Item ID"
// @Param        request body UpdateCartRequest true "Update Quantity Payload"
// @Success      200 {object} response.JSONResponse
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Router       /cart/{id} [put]
func (h *Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid cart item ID")
		return
	}

	var req UpdateCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := h.service.UpdateItem(userID, uint(id), req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "Cart item updated", nil)
}

// RemoveItem godoc
// @Summary      Remove item from cart
// @Description  Remove a single item from the user's shopping cart by item ID
// @Tags         cart
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Cart Item ID"
// @Success      200 {object} response.JSONResponse
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /cart/{id} [delete]
func (h *Handler) RemoveItem(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid cart item ID")
		return
	}
	if err := h.service.RemoveItem(userID, uint(id)); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to remove item")
		return
	}
	response.Success(w, http.StatusOK, "Cart item removed", nil)
}

// ClearCart godoc
// @Summary      Clear shopping cart
// @Description  Remove all items from the authenticated user's shopping cart
// @Tags         cart
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /cart [delete]
func (h *Handler) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)
	if err := h.service.ClearCart(userID); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to clear cart")
		return
	}
	response.Success(w, http.StatusOK, "Cart cleared", nil)
}
