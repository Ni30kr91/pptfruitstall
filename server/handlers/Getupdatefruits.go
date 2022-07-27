package handlers

import (
	"net/http"
	"pptfruitstall/pkg"

	"github.com/gin-gonic/gin"
)

func Getupdate_fruits(c *gin.Context) {
	reqBody := pkg.Fruits{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusBadGateway, res)
	}
	result := pkg.UpdateFruitsService(reqBody)

	if result == false {
		res := gin.H{
			"error": "Something went Wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"error": "Fruits Update Successfully",
	}
	c.JSON(http.StatusCreated, res)
	return
}
