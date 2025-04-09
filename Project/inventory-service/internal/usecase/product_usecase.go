package usecase

import (
	"Assignment1/Project/inventory-service/internal/domain"
	"errors"
)

type ProductUsecase interface {
	CreateProduct(p *domain.Product) error
	GetProduct(id string) (*domain.Product, error)
	UpdateProduct(p *domain.Product) error
	DeleteProduct(id string) error
	ListProducts() ([]*domain.Product, error)
}

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(r domain.ProductRepository) *productUsecase {
	return &productUsecase{repo: r}
}

func (uc *productUsecase) CreateProduct(p *domain.Product) error {
	if p.Name == "" {
		return errors.New("product name required")
	}
	return uc.repo.Create(p)
}

func (uc *productUsecase) GetProduct(id string) (*domain.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *productUsecase) UpdateProduct(p *domain.Product) error {
	return uc.repo.Update(p)
}

func (uc *productUsecase) DeleteProduct(id string) error {
	return uc.repo.Delete(id)
}

func (uc *productUsecase) ListProducts() ([]*domain.Product, error) {
	return uc.repo.ListAll()
}
