package domain

type Order struct {
	ID       string
	Products []OrderItem
	Status   string
}

type OrderItem struct {
	ProductID string
	Quantity  int
}

type OrderRepository interface {
	Create(o *Order) error
	GetByID(id string) (*Order, error)
	Update(o *Order) error
	ListAll() ([]*Order, error)
}
