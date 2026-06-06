package product

import (
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
// @Summary      Get all products
// @Description  Retrieve products with optional filtering by search term or category, with pagination
// @Tags         products
// @Produce      json
// @Param        search    query     string  false  "Search by product name/description"
// @Param        category  query     string  false  "Filter by category"
// @Param        page      query     int     false  "Page number (default: 1)"
// @Param        limit     query     int     false  "Page size (default: 10)"
// @Success      200 {object} response.PaginationResponse{data=[]ProductResponse}
// @Failure      500 {object} response.JSONResponse
// @Router       /products [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	category := r.URL.Query().Get("category")

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, total, err := h.service.GetAll(search, category, page, limit)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	response.Pagination(w, "Products fetched successfully", products, page, limit, total)
}

// GetByID godoc
// @Summary      Get product by ID
// @Description  Retrieve a single product by its unique database ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200 {object} response.JSONResponse{data=ProductResponse}
// @Failure      400 {object} response.JSONResponse
// @Failure      404 {object} response.JSONResponse
// @Router       /products/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		response.Error(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.service.GetByID(uint(id))
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Product fetched successfully", product)
}
