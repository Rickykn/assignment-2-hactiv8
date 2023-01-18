package handlers

import (
	"fmt"
	"net/http"

	"github.com/Rickykn/assignment-2-hactiv8.git/controllers"
	"github.com/Rickykn/assignment-2-hactiv8.git/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var orderInput *models.OrderRequestDTO

	err := c.ShouldBindJSON(&orderInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	result := controllers.CreateOrder(orderInput)

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Success Create Order",
		"status code": http.StatusCreated,
		"data":        result,
	})

}

func GetAllOrder(c *gin.Context) {

	result := controllers.GetOrder()

	fmt.Println(result)

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Success Create Order",
		"status code": http.StatusCreated,
		"data":        result,
	})

}
