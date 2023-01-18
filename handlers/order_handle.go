package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Get Data Success",
		"status code": http.StatusCreated,
		"data":        result,
	})

}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	convId, _ := strconv.Atoi(id)

	i := controllers.DeleteOrder(convId)

	fmt.Println(i)

	if i == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":     "Failed Delete Order",
			"status code": http.StatusOK,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Success Delete Order",
		"status code": http.StatusOK,
	})

}
