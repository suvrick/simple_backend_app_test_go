package main

import (
	"kiss_web/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controllers.Init(r)
	r.Run(":8080")
}
