package response

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type PaginationMeta struct{
	Page int `json:"page"`
	Limt int `json:"limit"`
	Total int64 `json:"total"`
}

type PaginationResponse struct{
	Success bool `json:"success"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func Success(w http.ResponseWriter, status int, message string, data interface{}) {
	writeJSON(w, status, JSONResponse{
		Success: true,
		Message: message,
		Data: data,
	})
}

func Error(w http.ResponseWriter, status int, message string ) {
	writeJSON(w, status, JSONResponse{
		Success: false,
		Message: message,
		Data: nil,
	})
}

func Pagination(w http.ResponseWriter, message string, data interface{}, page, limit int, total int64) {
	writeJSON(w, http.StatusOK, PaginationResponse{
		Success: true,
		Message: message,
		Data: data,
		Meta: PaginationMeta{
			Page: page,
			Limt: limit,
			Total: total,
		},
	})
}