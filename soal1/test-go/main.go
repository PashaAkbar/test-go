package main

import (
	"github.com/PashaAkbar/test-go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*.html")

	r.GET("/", controllers.GetCountries)

	r.Run(":8080")
}
