package main

import (
	"log"
	"net/http"

	"github.com/havilz/ecommerce-backend/config"
	adminModule "github.com/havilz/ecommerce-backend/internal/modules/admin"
	authModule "github.com/havilz/ecommerce-backend/internal/modules/auth"
	cartModule "github.com/havilz/ecommerce-backend/internal/modules/cart"
	categoryModule "github.com/havilz/ecommerce-backend/internal/modules/category"
	orderModule "github.com/havilz/ecommerce-backend/internal/modules/order"
	productModule "github.com/havilz/ecommerce-backend/internal/modules/product"
	"github.com/havilz/ecommerce-backend/pkg/database"
	"github.com/havilz/ecommerce-backend/pkg/middleware"

	_ "github.com/havilz/ecommerce-backend/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title           E-Commerce API
// @version         1.0
// @description     API Server for E-Commerce Platform built with native Go, GORM, SQLite, and JWT.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@ecommerce.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Insert JWT token in format: Bearer <token>
func main() {
	// Load config
	cfg := config.Load()

	// Initialize database & auto-migrate
	db := database.InitDB(cfg.DBPath)

	// Set JWT secret
	middleware.SetJWTSecret(cfg.JWTSecret)

	// Repositories
	authRepo := authModule.NewRepository(db)
	productRepo := productModule.NewRepository(db)
	categoryRepo := categoryModule.NewRepository(db)
	cartRepo := cartModule.NewRepository(db)
	orderRepo := orderModule.NewRepository(db)
	adminRepo := adminModule.NewRepository(db)

	// Services
	authSvc := authModule.NewService(authRepo, cfg.JWTSecret, cfg.JWTExpiryHours)
	productSvc := productModule.NewService(productRepo)
	categorySvc := categoryModule.NewService(categoryRepo)
	cartSvc := cartModule.NewService(cartRepo, productRepo)
	orderSvc := orderModule.NewService(orderRepo, cartRepo, productRepo, db)
	adminSvc := adminModule.NewService(adminRepo)

	// Handlers
	authHandler := authModule.NewHandler(authSvc)
	productHandler := productModule.NewHandler(productSvc)
	categoryHandler := categoryModule.NewHandler(categorySvc)
	cartHandler := cartModule.NewHandler(cartSvc)
	orderHandler := orderModule.NewHandler(orderSvc)
	adminHandler := adminModule.NewHandler(adminSvc)

	// Router
	mux := http.NewServeMux()

	// Root Landing Route
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "E-Commerce API is running. Visit /swagger/index.html for API documentation."}`))
	})

	// Swagger Redirect Route
	mux.HandleFunc("GET /swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	// Swagger Route
	mux.Handle("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Auth (public)
	mux.HandleFunc("POST /api/v1/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/v1/auth/login", authHandler.Login)
	mux.Handle("GET /api/v1/auth/me",
		middleware.JWTMiddleware(http.HandlerFunc(authHandler.Me)))

	// Products (public)
	mux.HandleFunc("GET /api/v1/products", productHandler.GetAll)
	mux.HandleFunc("GET /api/v1/products/{id}", productHandler.GetByID)

	// Cart (JWT protected)
	mux.Handle("GET /api/v1/cart",
		middleware.JWTMiddleware(http.HandlerFunc(cartHandler.GetCart)))
	mux.Handle("POST /api/v1/cart",
		middleware.JWTMiddleware(http.HandlerFunc(cartHandler.AddItem)))
	mux.Handle("PUT /api/v1/cart/{id}",
		middleware.JWTMiddleware(http.HandlerFunc(cartHandler.UpdateItem)))
	mux.Handle("DELETE /api/v1/cart/{id}",
		middleware.JWTMiddleware(http.HandlerFunc(cartHandler.RemoveItem)))
	mux.Handle("DELETE /api/v1/cart",
		middleware.JWTMiddleware(http.HandlerFunc(cartHandler.ClearCart)))

	// Orders (JWT protected)
	mux.Handle("POST /api/v1/orders/checkout",
		middleware.JWTMiddleware(http.HandlerFunc(orderHandler.Checkout)))
	mux.Handle("GET /api/v1/orders",
		middleware.JWTMiddleware(http.HandlerFunc(orderHandler.GetOrders)))
	mux.Handle("GET /api/v1/orders/{id}",
		middleware.JWTMiddleware(http.HandlerFunc(orderHandler.GetOrderByID)))

	// Categories (public)
	mux.HandleFunc("GET /api/v1/categories", categoryHandler.GetAll)
	mux.HandleFunc("GET /api/v1/categories/{id}", categoryHandler.GetByID)

	// Admin (JWT + AdminOnly)
	adminMW := func(h http.Handler) http.Handler {
		return middleware.JWTMiddleware(middleware.AdminOnly(h))
	}
	mux.Handle("POST /api/v1/admin/products",
		adminMW(http.HandlerFunc(adminHandler.CreateProduct)))
	mux.Handle("PUT /api/v1/admin/products/{id}",
		adminMW(http.HandlerFunc(adminHandler.UpdateProduct)))
	mux.Handle("DELETE /api/v1/admin/products/{id}",
		adminMW(http.HandlerFunc(adminHandler.DeleteProduct)))
	mux.Handle("GET /api/v1/admin/orders",
		adminMW(http.HandlerFunc(adminHandler.GetAllOrders)))
	mux.Handle("PATCH /api/v1/admin/orders/{id}",
		adminMW(http.HandlerFunc(adminHandler.UpdateOrderStatus)))

	// Categories Admin (AdminOnly)
	mux.Handle("POST /api/v1/admin/categories",
		adminMW(http.HandlerFunc(categoryHandler.Create)))
	mux.Handle("PUT /api/v1/admin/categories/{id}",
		adminMW(http.HandlerFunc(categoryHandler.Update)))
	mux.Handle("DELETE /api/v1/admin/categories/{id}",
		adminMW(http.HandlerFunc(categoryHandler.Delete)))

	//Start server
	addr := ":" + cfg.Port
	log.Printf("[server] running on http://localhost%s", addr)
	corsHandler := middleware.CORSMiddleware(cfg.AllowedOrigins)
	if err := http.ListenAndServe(addr, corsHandler(mux)); err != nil {
		log.Fatalf("[server] failed to start: %v", err)
	}
}
