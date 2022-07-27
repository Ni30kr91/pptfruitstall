package server

import (
	"github.com/gin-gonic/gin"
)

func Server() {

	r := gin.Default()
	setupRoutes(r)
	r.Run()
}
