package service

import (
	"context"
	"fmt"
	"go-grpc/commons/models"
	block "go-grpc/commons/pb"
	"go-grpc/service/database"
	"log"
	"time"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const (
	ProjectId               = "athena-dsv"
	precipitationDataTypeId = "6"
	temperatureDataTypeId   = "1"
	tableName               = ProjectId + ".athena.pixel"
)

// GetBlockById BlocksServer interface implementation, consults the pixel table on BigQuery
//
// retrieving the last data of the day 2022-06-28 matching the id given in the request message arg and
// returns the responseBlock message
func (server *Server) GetBlockById(context context.Context, req *block.RequestID) (*block.ResponseBlock, error) {
	bqQuery := server.BQClient.Query(getQueryBlock())
	bqQuery.Parameters = []bigquery.QueryParameter{
		{
			Name:  "source_id",
			Value: req.GetId(),
		},
	}
	iter, err := database.QueryConnection(bqQuery, context)
	if err != nil {
		log.Fatalln("error in query connection:", err)
		return nil, err
	}

	var blockDAO models.Block
	for {
		var row []bigquery.Value
		err := iter.Next(&row)
		if err == iterator.Done {
			fmt.Println("Iterator done", err)
			break
		}
		if err != nil {
			log.Println("error in iter.Next", err)
			return nil, err
		}
		blockDAO.DataTimestamp = row[0].(time.Time)
		blockDAO.CreatedTimestamp = row[1].(time.Time)
		blockDAO.TemperatureInst = row[2].(string)
		blockDAO.TemperatureMin = row[3].(string)
		blockDAO.TemperatureMax = row[4].(string)
		blockDAO.PrecipitationInst = row[5].(string)
		fmt.Println(blockDAO)
	}

	resBlock := models.TransformBlockDAOIntoResponse(blockDAO)
	return resBlock, nil
}

func getQueryBlock() string {
	return "select (a_temp.data_timestamp) as data_timestamp, (a_temp.created_timestamp) as created_timestamp, " +
		" a_temp.data_inst as temp_inst, a_temp.data_min as temp_min, a_temp.data_max as temp_max, " +
		" a_prec.data_inst as prec_inst " +
		" from `" + tableName + "` a_temp " +
		" inner join ( " +
		" select data_timestamp, created_timestamp, data_inst, data_min, data_max, source_id " +
		" from `" + tableName + "` " +
		` where date(created_timestamp) = "2022-06-28" ` +
		" and datatype_id = " + precipitationDataTypeId +
		" and source_type_id = 3 " +
		" and source_subtype_id = 16 " +
		" and source_id = @source_id " +
		" ) as a_prec " +
		" on a_temp.source_id = a_prec.source_id " +
		` where date(a_temp.created_timestamp) = "2022-06-28" ` +
		" and a_temp.datatype_id = " + temperatureDataTypeId +
		" and a_temp.source_type_id = 3 " +
		" and a_temp.source_subtype_id = 16 " +
		" and a_temp.source_id = @source_id " +
		" order by a_temp.created_timestamp desc, a_prec.created_timestamp desc " +
		" limit 1 "
}
