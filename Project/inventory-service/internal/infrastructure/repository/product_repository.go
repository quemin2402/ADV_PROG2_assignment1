package repository

import (
	"Assignment1/Project/inventory-service/internal/domain"
	"database/sql"
	"errors"
)

type postgresProductRepo struct {
	db *sql.DB
}

func NewPostgresProductRepo(db *sql.DB) *postgresProductRepo {
	return &postgresProductRepo{db: db}
}

func (r *postgresProductRepo) Create(p *domain.Product) error {
	query := `INSERT INTO products (id, name, category, price, stock)
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, p.ID, p.Name, p.Category, p.Price, p.Stock)
	if err != nil {
		return err
	}
	return nil
}

func (r *postgresProductRepo) GetByID(id string) (*domain.Product, error) {
	query := `SELECT id, name, category, price, stock 
              FROM products WHERE id = $1 LIMIT 1`
	row := r.db.QueryRow(query, id)

	var prod domain.Product
	if err := row.Scan(&prod.ID, &prod.Name, &prod.Category, &prod.Price, &prod.Stock); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &prod, nil
}

func (r *postgresProductRepo) Update(p *domain.Product) error {
	_, err := r.GetByID(p.ID)
	if err != nil {
		return err
	}

	query := `UPDATE products
              SET name = $2, category = $3, price = $4, stock = $5
              WHERE id = $1`
	_, err = r.db.Exec(query, p.ID, p.Name, p.Category, p.Price, p.Stock)
	return err
}

func (r *postgresProductRepo) Delete(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (r *postgresProductRepo) ListAll() ([]*domain.Product, error) {
	query := `SELECT id, name, category, price, stock FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var prod domain.Product
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Category, &prod.Price, &prod.Stock); err != nil {
			return nil, err
		}
		products = append(products, &prod)
	}
	return products, rows.Err()
}
