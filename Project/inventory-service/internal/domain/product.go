package domain

type Product struct {
	ID       string
	Name     string
	Category string
	Price    float64
	Stock    int
}

type ProductRepository interface {
	Create(p *Product) error
	GetByID(id string) (*Product, error)
	Update(p *Product) error
	Delete(id string) error
	ListAll() ([]*Product, error)
}
