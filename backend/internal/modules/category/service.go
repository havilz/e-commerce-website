package category

import (
	"errors"
	"strings"

	"github.com/havilz/ecommerce-backend/models"
	"github.com/havilz/ecommerce-backend/pkg/security"
)

type Service interface {
	Create(req CreateCategoryRequest) (*CategoryResponse, error)
	Update(id uint, req UpdateCategoryRequest) (*CategoryResponse, error)
	Delete(id uint) error
	GetAll() ([]CategoryResponse, error)
	GetByID(id uint) (*CategoryResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(req CreateCategoryRequest) (*CategoryResponse, error) {
	name := strings.TrimSpace(req.Name)
	name = security.SanitizeString(name)
	if name == "" {
		return nil, errors.New("category name is required")
	}

	// Check if category already exists
	existing, err := s.repo.FindByName(name)
	if err == nil && existing != nil {
		return nil, errors.New("category already exists")
	}

	cat := &models.Category{Name: name}
	if err := s.repo.Create(cat); err != nil {
		return nil, errors.New("failed to create category")
	}

	return &CategoryResponse{
		ID:        cat.ID,
		Name:      cat.Name,
		CreatedAt: cat.CreatedAt,
		UpdatedAt: cat.UpdatedAt,
	}, nil
}

func (s *service) Update(id uint, req UpdateCategoryRequest) (*CategoryResponse, error) {
	name := strings.TrimSpace(req.Name)
	name = security.SanitizeString(name)
	if name == "" {
		return nil, errors.New("category name is required")
	}

	cat, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	// Check if duplicate name with another category
	existing, err := s.repo.FindByName(name)
	if err == nil && existing != nil && existing.ID != id {
		return nil, errors.New("category name already exists")
	}

	cat.Name = name
	if err := s.repo.Update(cat); err != nil {
		return nil, errors.New("failed to update category")
	}

	return &CategoryResponse{
		ID:        cat.ID,
		Name:      cat.Name,
		CreatedAt: cat.CreatedAt,
		UpdatedAt: cat.UpdatedAt,
	}, nil
}

func (s *service) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}
	return s.repo.Delete(id)
}

func (s *service) GetAll() ([]CategoryResponse, error) {
	cats, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []CategoryResponse
	for _, c := range cats {
		result = append(result, CategoryResponse{
			ID:        c.ID,
			Name:      c.Name,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
	if result == nil {
		result = []CategoryResponse{}
	}
	return result, nil
}

func (s *service) GetByID(id uint) (*CategoryResponse, error) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}
	return &CategoryResponse{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}, nil
}
