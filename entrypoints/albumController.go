package entrypoints

import (
	"net/http"

	"github.com/dhfherna/microservice-go-poc/services"
	"github.com/gin-gonic/gin"
)

func getAlbumsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"albums": services.GetAlbums(),
	})
}

func AlbumsController(router *gin.Engine) {
	router.GET("/albums", getAlbumsHandler)
}
