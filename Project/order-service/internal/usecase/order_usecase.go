package usecase

import (
	"Assignment1/Project/order-service/internal/domain"
	"errors"
)

type OrderUsecase interface {
	CreateOrder(o *domain.Order) error
	GetOrder(id string) (*domain.Order, error)
	UpdateOrder(o *domain.Order) error
	ListOrders() ([]*domain.Order, error)
}

type orderUsecase struct {
	repo domain.OrderRepository
}

func NewOrderUsecase(r domain.OrderRepository) OrderUsecase {
	return &orderUsecase{repo: r}
}

func (uc *orderUsecase) CreateOrder(o *domain.Order) error {
	if len(o.Products) == 0 {
		return errors.New("order must contain at least one product")
	}
	o.Status = "pending"
	return uc.repo.Create(o)
}

func (uc *orderUsecase) GetOrder(id string) (*domain.Order, error) {
	return uc.repo.GetByID(id)
}

func (uc *orderUsecase) UpdateOrder(o *domain.Order) error {
	return uc.repo.Update(o)
}

func (uc *orderUsecase) ListOrders() ([]*domain.Order, error) {
	return uc.repo.ListAll()
}
