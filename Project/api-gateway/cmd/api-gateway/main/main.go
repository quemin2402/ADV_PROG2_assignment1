package main

import (
	"Assignment1/Project/api-gateway/internal/gateway"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	gateway.SetupRoutes(r)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
