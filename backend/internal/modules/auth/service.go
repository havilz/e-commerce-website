package auth

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/havilz/ecommerce-backend/models"
	"github.com/havilz/ecommerce-backend/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type Service interface {
	Register(req RegisterRequest) (*AuthResponse, error)
	Login(req LoginRequest) (*AuthResponse, error)
	GetMe(userID uint) (*MeResponse, error)
}

type service struct {
	repo        Repository
	jwtSecret   string
	expiryHours int
}

func NewService(repo Repository, jwtSecret string, expiryHours int) Service {
	return &service{
		repo:        repo,
		jwtSecret:   jwtSecret,
		expiryHours: expiryHours,
	}
}

func (s *service) Register(req RegisterRequest) (*AuthResponse, error) {
	// Sanitize text inputs to strip out HTML/JS tags
	req.Name = security.SanitizeString(req.Name)
	req.Email = security.SanitizeString(req.Email)

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("name, email, and password are required")
	}

	if !emailRegex.MatchString(req.Email) {
		return nil, errors.New("invalid email format")
	}

	if len(req.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	existing, _ := s.repo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Role:     "user",
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  UserData{ID: user.ID, Name: user.Name, Email: user.Email, Role: user.Role},
	}, nil
}

func (s *service) Login(req LoginRequest) (*AuthResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  UserData{ID: user.ID, Name: user.Name, Email: user.Email, Role: user.Role},
	}, nil
}

func (s *service) GetMe(userID uint) (*MeResponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &MeResponse{ID: user.ID, Name: user.Name, Email: user.Email, Role: user.Role}, nil
}

func (s *service) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Duration(s.expiryHours) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signed, nil
}
