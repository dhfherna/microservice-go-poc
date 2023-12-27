package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	host := flag.String("host", "localhost", "Host de la aplicación")
	port := flag.String("port", "8080", "Puerto de la aplicación")
	flag.Parse()

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to my microservice!",
		})
	})

	router.Run(fmt.Sprintf("%s:%s", *host, *port))
}
