package product

import "errors"

type Service interface {
	GetAll(search, category string, page, limit int) ([]ProductResponse, int64, error)
	GetByID(id uint) (*ProductResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll(search, category string, page, limit int) ([]ProductResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	products, total, err := s.repo.FindAll(search, category, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var result []ProductResponse
	for _, p := range products {
		result = append(result, ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
			ImageURL:    p.ImageURL,
			Category:    p.Category,
		})
	}
	return result, total, nil
}

func (s *service) GetByID(id uint) (*ProductResponse, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return &ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		ImageURL:    p.ImageURL,
		Category:    p.Category,
	}, nil
}
