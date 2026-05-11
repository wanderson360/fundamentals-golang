package controller

import (
	usecase "api-product/useCase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return ProductController{ProductUseCase: usecase}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.ProductUseCase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
