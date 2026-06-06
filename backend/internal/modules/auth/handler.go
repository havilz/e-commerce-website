package auth

import (
	"encoding/json"
	"net/http"

	"github.com/havilz/ecommerce-backend/pkg/middleware"
	"github.com/havilz/ecommerce-backend/pkg/response"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account with default role "user"
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body RegisterRequest true "Register Payload"
// @Success      201 {object} response.JSONResponse{data=MeResponse}
// @Failure      400 {object} response.JSONResponse
// @Router       /auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.Register(req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, "User registered successfully", result)
}

// Login godoc
// @Summary      User Login
// @Description  Authenticate user credentials and return a JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body LoginRequest true "Login Payload"
// @Success      200 {object} response.JSONResponse{data=AuthResponse}
// @Failure      400 {object} response.JSONResponse
// @Failure      401 {object} response.JSONResponse
// @Router       /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.Login(req)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "Login successful", result)
}

// Me godoc
// @Summary      Get current user profile
// @Description  Retrieve the profile info of the currently authenticated user
// @Tags         auth
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.JSONResponse{data=MeResponse}
// @Failure      401 {object} response.JSONResponse
// @Failure      404 {object} response.JSONResponse
// @Router       /auth/me [get]
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r)

	result, err := h.service.GetMe(userID)
	if err != nil {
		response.Error(w, http.StatusNotFound, err.Error())
		return
	}

	response.Success(w, http.StatusOK, "User fetched successfully", result)
}
