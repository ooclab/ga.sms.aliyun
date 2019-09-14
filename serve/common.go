package serve

import (
	"github.com/gin-gonic/gin"
)

func (api *API) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (api *API) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "success",
	})
}
