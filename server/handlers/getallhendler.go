package handlers

import (
	"net/http"
	"pptfruitstall/pkg"

	"github.com/gin-gonic/gin"
)

func Getallfruits(c *gin.Context) {

	data := pkg.GetFruitsService()

	res := gin.H{
		"data": data,
	}
	c.JSON(http.StatusOK, res)
	return

}
