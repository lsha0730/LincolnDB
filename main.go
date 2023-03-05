package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "io/ioutil"
	// "os"
	// "path/filepath"
	// "sync"
	// "github.com/jcelliott/lumber"
	"github.com/lsha0730/LycheeDB/util"
)

var PORT string = ":8000"
var ROOT string = "./"

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.POST("/", func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		// TODO: Add tolerance for simple JSON query
		queryString := data["query"].(string)
		query, parseError := util.StrToMap(queryString)
		if parseError != nil {
			c.AbortWithStatusJSON(http.StatusOK, parseError.Error())
			return
		}
		if err := handleQuery(query); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, err.Error())
			return
		}

		c.AbortWithStatusJSON(http.StatusOK, "success")
		return
	})

	r.Run(PORT)
}

func handleQuery(query map[string]interface{}) error {
	if err := util.ValidateQuery(query); err != nil {
		return err
	}

	// TODO: Actually do something with the query
	db := query["db"].(string)
	op := query["op"].(string)
	path := query["path"].(string)

	driver, _ := util.NewDB(db)

	switch op {
	case util.READ:
		driver.HandleRead(path)
	case util.WRITE:
		driver.HandleWrite(path, query["value"])
	case util.LIST:
		driver.HandleList(path)
	case util.MAKEDB:
		util.NewDB(query["name"].(string))
	default:
		return errors.New("ERROR: Unexpected operation type")
	}

	return nil
}
