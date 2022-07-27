package server

import (
	"pptfruitstall/server/handlers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {

	r.GET("/fruits", handlers.Getallfruits)
	r.POST("/add_fruits", handlers.Getadd_fruits)
	r.DELETE("/delete_fruits/:id", handlers.Getdelete_fruits)
	r.PUT("/update_fruits", handlers.Getupdate_fruits)
}
