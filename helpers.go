package helpers

import "net/http"

func getPage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
