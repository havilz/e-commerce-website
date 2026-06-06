package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/havilz/ecommerce-backend/pkg/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Add a new product to the catalog (Admin Only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body AdminProductRequest true "Create Product Payload"
// @Success      201 {object} response.JSONResponse{data=AdminProductResponse}
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      403 {object} response.JSONResponse
// @Router       /admin/products [post]
func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req AdminProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	p, err := h.service.CreateProduct(req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, http.StatusCreated, "Product created successfully", p)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update the details of an existing product (Admin Only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Product ID"
// @Param        request body AdminProductRequest true "Update Product Payload"
// @Success      200 {object} response.JSONResponse
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      403 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /admin/products/{id} [put]
func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	var req AdminProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := h.service.UpdateProduct(uint(id), req); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to update product")
		return
	}
	response.Success(w, http.StatusOK, "Product updated successfully", nil)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product from the database (Admin Only)
// @Tags         admin
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Product ID"
// @Success      200 {object} response.JSONResponse
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      403 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /admin/products/{id} [delete]
func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	if err := h.service.DeleteProduct(uint(id)); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}
	response.Success(w, http.StatusOK, "Product deleted successfully", nil)
}

// GetAllOrders godoc
// @Summary      Get all orders
// @Description  Retrieve a paginated list of all customer orders in the system (Admin Only)
// @Tags         admin
// @Produce      json
// @Security     BearerAuth
// @Param        page   query     int  false  "Page number (default: 1)"
// @Param        limit  query     int  false  "Page size (default: 10)"
// @Success      200 {object} response.PaginationResponse{data=[]AdminOrderResponse}
// @Failure      401 {object} response.JSONResponse
// @Failure      403 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /admin/orders [get]
func (h *Handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	orders, total, err := h.service.GetAllOrders(page, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch orders")
		return
	}
	response.Pagination(w, "Orders fetched successfully", orders, page, limit, total)
}

// UpdateOrderStatus godoc
// @Summary      Update order status
// @Description  Update the delivery status of a client's order (Admin Only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Order ID"
// @Param        request body UpdateOrderStatusRequest true "Update Status Payload"
// @Success      200 {object} response.JSONResponse
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      403 {object} response.JSONResponse
// @Router       /admin/orders/{id} [patch]
func (h *Handler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid order ID")
		return
	}
	var req UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if err := h.service.UpdateOrderStatus(uint(id), req); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "Order status updated successfully", nil)
}
