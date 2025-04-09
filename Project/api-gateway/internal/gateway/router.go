package gateway

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.GET("/gateway/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "API Gateway OK"})
	})

	r.Any("/products/*proxyPath", InventoryProxy)
	r.Any("/orders/*proxyPath", OrderProxy)
}
