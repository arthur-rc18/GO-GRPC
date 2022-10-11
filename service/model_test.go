package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go-grpc/commons/models"
	block "go-grpc/commons/pb"
	"go-grpc/commons/testUtils"
	"testing"
)

var sourceId = "B233"
var ctx = context.Background()

func TestGetBlockById(t *testing.T) {
	client := testUtils.SetupGrpcTestClient()
	t.Run("get existent block", func(t *testing.T) {
		t.Parallel()
		expect := models.TransformBlockDAOIntoResponse(models.Block{})
		got, _ := client.GetBlockById(ctx, &block.RequestID{Id: sourceId})
		assert.NotEqual(t, expect.GetBlockJson(), got.GetBlockJson())
	})

	t.Run("get unexistant block", func(t *testing.T) {
		t.Parallel()
		expect := models.TransformBlockDAOIntoResponse(models.Block{})
		got, _ := client.GetBlockById(ctx, &block.RequestID{Id: "asdfasfd"})
		assert.Equal(t, expect.GetBlockJson(), got.GetBlockJson())
	})
}
