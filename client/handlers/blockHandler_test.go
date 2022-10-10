package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-grpc/commons"
	"go-grpc/commons/models"
	block "go-grpc/commons/pb"
	"go-grpc/commons/testUtils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var sourceId = "B233"

func setupTestingRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r
}

func TestGetBlockById(t *testing.T) {

	r := setupTestingRoutes()
	r.GET(*commons.ROUTEApi+":id", GetBlockByIdHandler)

	client := testUtils.SetupGrpcTestClient()

	t.Run("Getting existent block", func(t *testing.T) {
		t.Parallel()

		req, _ := http.NewRequest("GET", *commons.ROUTEApi+sourceId, nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		var gotBlock models.Block
		_ = json.Unmarshal(res.Body.Bytes(), &gotBlock)
		fmt.Println(&gotBlock)
		assert.Equal(t, http.StatusOK, 200)

		resBlock, _ := client.GetBlockById(context.Background(), &block.RequestID{Id: sourceId})
		expect := models.TransformResponseIntoBlockDTO(resBlock)
		assert.Equal(t, expect, gotBlock)
	})

	t.Run("Getting non existent block", func(t *testing.T) {
		t.Parallel()

		req, _ := http.NewRequest("GET", *commons.ROUTEApi+"asdfasdf", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
