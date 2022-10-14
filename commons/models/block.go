package models

import (
	"encoding/json"
	"fmt"
	block "go-grpc/commons/pb"
	"time"
)

// Block defines the block data structure to be serialized and deserialized
type Block struct {
	DataTimestamp     time.Time `json:"data_timestamp"`
	CreatedTimestamp  time.Time `json:"created_timestamp"`
	TemperatureInst   string    `json:"temperature_inst"`
	TemperatureMin    string    `json:"temperature_min"`
	TemperatureMax    string    `json:"temperature_max"`
	PrecipitationInst string    `json:"precipitation_inst"`
}

// TransformBlockDAOIntoResponse serializes a block to json/string format to compose the grpc responseBlock message
func TransformBlockDAOIntoResponse(blockDAO Block) *block.ResponseBlock {
	var responseBlock block.ResponseBlock
	blockBytes, err := json.Marshal(blockDAO)
	if err != nil {
		fmt.Println("error marshalling block", err)
		return nil
	}
	responseBlock.BlockJson = string(blockBytes)
	return &responseBlock
}

// TransformResponseIntoBlockDTO deserializes the responseBlock message as json/string format to Block
func TransformResponseIntoBlockDTO(response *block.ResponseBlock) Block {
	var resBlock Block
	err := json.Unmarshal([]byte(response.BlockJson), &resBlock)
	if err != nil {
		fmt.Println("error unmarshalling block", err)
		return Block{}
	}
	return resBlock
}
