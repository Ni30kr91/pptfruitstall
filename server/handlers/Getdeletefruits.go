package handlers

import (
	"net/http"
	"pptfruitstall/pkg"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Getdelete_fruits(c *gin.Context) {
	id, err := c.Params.Get("id")
	if err == false {
		res := gin.H{
			"error": "something went wrong",
		}
		c.JSON(http.StatusBadGateway, res)
	}
	id_int, _ := strconv.Atoi(id)
	result := pkg.DeleteFruitsService(id_int)

	if result == false {
		res := gin.H{
			"error": "something went wrong",
		}
		c.JSON(http.StatusBadGateway, res)
		return
	}
	res := gin.H{
		"error": "Delete successfully",
	}
	c.JSON(http.StatusOK, res)
	return
}
