package category

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

// GetAll godoc
// @Summary      Get all categories
// @Description  Retrieve a list of all product categories
// @Tags         categories
// @Produce      json
// @Success      200 {object} response.JSONResponse{data=[]CategoryResponse}
// @Failure      500 {object} response.JSONResponse
// @Router       /categories [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	cats, err := h.service.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}
	response.Success(w, http.StatusOK, "Categories fetched successfully", cats)
}

// GetByID godoc
// @Summary      Get category by ID
// @Description  Retrieve a single category by its unique database ID
// @Tags         categories
// @Produce      json
// @Param        id   path      int  true  "Category ID"
// @Success      200 {object} response.JSONResponse{data=CategoryResponse}
// @Failure      400 {object} response.JSONResponse
// @Failure      404 {object} response.JSONResponse
// @Router       /categories/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	cat, err := h.service.GetByID(uint(id))
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Category fetched successfully", cat)
}

// Create godoc
// @Summary      Create a new category
// @Description  Add a new category (Admin Only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body CreateCategoryRequest true "Create Category Payload"
// @Success      201 {object} response.JSONResponse{data=CategoryResponse}
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      403 {object} response.JSONResponse
// @Router       /admin/categories [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	cat, err := h.service.Create(req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, http.StatusCreated, "Category created successfully", cat)
}

// Update godoc
// @Summary      Update a category
// @Description  Update details of an existing category (Admin Only)
// @Tags         admin
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Category ID"
// @Param        request body UpdateCategoryRequest true "Update Category Payload"
// @Success      200 {object} response.JSONResponse{data=CategoryResponse}
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      403 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /admin/categories/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	var req UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	cat, err := h.service.Update(uint(id), req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "Category updated successfully", cat)
}

// Delete godoc
// @Summary      Delete a category
// @Description  Delete a category (Admin Only)
// @Tags         admin
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "Category ID"
// @Success      200 {object} response.JSONResponse
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Failure      403 {object} response.JSONResponse
// @Failure      500 {object} response.JSONResponse
// @Router       /admin/categories/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(w, http.StatusOK, "Category deleted successfully", nil)
}
