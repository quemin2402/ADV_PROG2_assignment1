package main

import (
	"Assignment1/Project/inventory-service/internal/handler"
	_ "Assignment1/Project/inventory-service/internal/handler"
	"Assignment1/Project/inventory-service/internal/infrastructure/db"
	"Assignment1/Project/inventory-service/internal/infrastructure/repository"
	"Assignment1/Project/inventory-service/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	pgDB, err := db.NewPostgresDB("localhost", "5432", "postgres", "0000", "inventory_service")
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	_, err = pgDB.Exec(`
        CREATE TABLE IF NOT EXISTS products (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL,
            category TEXT NOT NULL,
            price NUMERIC(12,2) NOT NULL,
            stock INT NOT NULL
        );
    `)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	productRepo := repository.NewPostgresProductRepo(pgDB)
	productUC := usecase.NewProductUsecase(productRepo)

	productHandler := handler.NewProductHandler(productUC)

	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products/:id", productHandler.GetProduct)
	r.PATCH("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)
	r.GET("/products", productHandler.ListProducts)

	log.Println("Inventory Service running on :8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}
