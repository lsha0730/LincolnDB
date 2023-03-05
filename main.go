package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// "io/ioutil"
	// "os"
	// "path/filepath"
	// "sync"
	// "github.com/jcelliott/lumber"
	"github.com/lsha0730/LycheeDB/util"
)

var PORT string = ":8000"

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var data map[string]string
		if err := c.ShouldBindJSON(&data); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		queryString := data["query"]
		if err := handleQuery(queryString); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, err.Error())
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, "success")
	})

	r.Run(PORT)
}

func handleQuery(queryString string) error {
	if err := util.ValidateQuery(queryString); err != nil {
		return err
	}

	// TODO: Actually do something with the query

	return nil
}
