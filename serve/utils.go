package serve

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) writeError(c *gin.Context, data interface{}) {
	b, _ := json.MarshalIndent(data, "", "    ")
	fmt.Println(string(b))
	c.JSON(http.StatusBadRequest, data)
}
