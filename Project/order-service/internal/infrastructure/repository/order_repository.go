package repository

import (
	"Assignment1/Project/order-service/internal/domain"
	"database/sql"
	"errors"
)

type postgresOrderRepo struct {
	db *sql.DB
}

func NewPostgresOrderRepo(db *sql.DB) domain.OrderRepository {
	return &postgresOrderRepo{db: db}
}

func (r *postgresOrderRepo) Create(o *domain.Order) error {
	insertOrderSQL := `INSERT INTO orders (id, status) VALUES ($1, $2)`
	_, err := r.db.Exec(insertOrderSQL, o.ID, o.Status)
	if err != nil {
		return err
	}

	insertItemSQL := `INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`
	for _, item := range o.Products {
		_, err = r.db.Exec(insertItemSQL, o.ID, item.ProductID, item.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *postgresOrderRepo) GetByID(id string) (*domain.Order, error) {
	orderQuery := `SELECT id, status FROM orders WHERE id = $1 LIMIT 1`
	row := r.db.QueryRow(orderQuery, id)

	var ord domain.Order
	if err := row.Scan(&ord.ID, &ord.Status); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}

	itemsQuery := `SELECT product_id, quantity FROM order_items WHERE order_id = $1`
	rows, err := r.db.Query(itemsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.OrderItem
	for rows.Next() {
		var item domain.OrderItem
		if err := rows.Scan(&item.ProductID, &item.Quantity); err != nil {
			return nil, err
		}
		products = append(products, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	ord.Products = products
	return &ord, nil
}

func (r *postgresOrderRepo) Update(o *domain.Order) error {
	existing, err := r.GetByID(o.ID)
	if err != nil {
		return err
	}
	existing.Status = o.Status
	existing.Products = o.Products

	updateOrderSQL := `UPDATE orders SET status = $2 WHERE id = $1`
	_, err = r.db.Exec(updateOrderSQL, existing.ID, existing.Status)
	if err != nil {
		return err
	}

	deleteItemsSQL := `DELETE FROM order_items WHERE order_id = $1`
	_, err = r.db.Exec(deleteItemsSQL, existing.ID)
	if err != nil {
		return err
	}

	insertItemSQL := `INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`
	for _, item := range existing.Products {
		_, err = r.db.Exec(insertItemSQL, existing.ID, item.ProductID, item.Quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *postgresOrderRepo) ListAll() ([]*domain.Order, error) {
	rows, err := r.db.Query(`SELECT id, status FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []*domain.Order
	for rows.Next() {
		var ord domain.Order
		if err := rows.Scan(&ord.ID, &ord.Status); err != nil {
			return nil, err
		}

		itemRows, err := r.db.Query(`SELECT product_id, quantity FROM order_items WHERE order_id=$1`, ord.ID)
		if err != nil {
			return nil, err
		}
		var items []domain.OrderItem
		for itemRows.Next() {
			var it domain.OrderItem
			if err := itemRows.Scan(&it.ProductID, &it.Quantity); err != nil {
				itemRows.Close()
				return nil, err
			}
			items = append(items, it)
		}
		itemRows.Close()

		ord.Products = items
		orders = append(orders, &ord)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
