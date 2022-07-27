package handlers

import (
	"net/http"
	"pptfruitstall/pkg"

	"github.com/gin-gonic/gin"
)

func Getadd_fruits(c *gin.Context) {
	reqBody := pkg.Fruits{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": "Something went Wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result := pkg.FruitInsertService(reqBody)

	if result == false {
		res := gin.H{
			"error": "Something went Wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"message": "fruits Inserted",
	}
	c.JSON(http.StatusCreated, res)
}
