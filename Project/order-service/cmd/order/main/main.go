package main

import (
	"Assignment1/Project/order-service/internal/handler"
	"Assignment1/Project/order-service/internal/infrastructure/db"
	pgRepo "Assignment1/Project/order-service/internal/infrastructure/repository"
	"Assignment1/Project/order-service/internal/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	pgDB, err := db.NewPostgresDB("localhost", "5432", "postgres", "0000", "order_service")
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	_, err = pgDB.Exec(`
        CREATE TABLE IF NOT EXISTS orders (
            id TEXT PRIMARY KEY,
            status TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatalf("Failed to create orders table: %v", err)
	}

	_, err = pgDB.Exec(`
        CREATE TABLE IF NOT EXISTS order_items (
            order_id TEXT NOT NULL,
            product_id TEXT NOT NULL,
            quantity INT NOT NULL,
            PRIMARY KEY (order_id, product_id),
            FOREIGN KEY (order_id) REFERENCES orders (id)
        );
    `)
	if err != nil {
		log.Fatalf("Failed to create order_items table: %v", err)
	}

	orderRepo := pgRepo.NewPostgresOrderRepo(pgDB)
	orderUC := usecase.NewOrderUsecase(orderRepo)

	orderHandler := handler.NewOrderHandler(orderUC)

	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders/:id", orderHandler.GetOrder)
	r.PATCH("/orders/:id", orderHandler.UpdateOrder)
	r.GET("/orders", orderHandler.ListOrders)

	log.Println("Order Service running on :8082")
	if err := http.ListenAndServe(":8082", r); err != nil {
		log.Fatal(err)
	}
}
