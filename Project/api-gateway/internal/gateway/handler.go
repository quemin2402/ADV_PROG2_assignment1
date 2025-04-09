package gateway

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const inventoryBaseURL = "http://localhost:8081"
const orderBaseURL = "http://localhost:8082"

func InventoryProxy(c *gin.Context) {
	proxyPath := c.Param("proxyPath")
	targetURL := fmt.Sprintf("%s/products%s", inventoryBaseURL, proxyPath)
	proxyRequest(c, targetURL)
}

func OrderProxy(c *gin.Context) {
	proxyPath := c.Param("proxyPath")
	targetURL := fmt.Sprintf("%s/orders%s", orderBaseURL, proxyPath)
	proxyRequest(c, targetURL)
}

func proxyRequest(c *gin.Context, targetURL string) {
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), bodyBytes)
}
