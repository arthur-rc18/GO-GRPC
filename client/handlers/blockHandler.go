// Package handlers is responsible for provide the application handlers(controllers)
package handlers

import (
	"go-grpc/client/service"
	"go-grpc/commons/models"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// GetBlockByIdHandler handles the route to get the last data of day 2022-06-28 from the block passed on id
//			GET /api/id
// if there is no block with the given source_id (id on route) returns null body and NotFound Status
//
// if the block exists on pixel table returns its data and Ok Status
func GetBlockByIdHandler(c *gin.Context) {
	id := c.Param("id")

	blockDTO, err := service.GetBlockById(id)

	if err != nil {
		log.Println("error block returning:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else if reflect.DeepEqual(blockDTO, models.Block{}) {
		log.Println("empty block:", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, blockDTO)
}
