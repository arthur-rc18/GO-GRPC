package database

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
)

// NewBQClient creates a new BigQuery client
func NewBQClient() (*bigquery.Client, error) {
	return bigquery.NewClient(context.Background(), "athena-dsv")
}

// QueryConnection execute a given *biquery.Query and return its iterator and an error
func QueryConnection(bqQuery *bigquery.Query, ctx context.Context) (*bigquery.RowIterator, error) {
	bqQuery.Location = "US"
	job, err := bqQuery.Run(ctx)
	if err != nil {
		log.Println("error query run - job creation:", err.Error())
		return nil, err
	}
	status, err := job.Wait(ctx)
	if err != nil {
		log.Println("error job wait - status", err.Error())
		return nil, err
	}
	if err := status.Err(); err != nil {
		log.Println("error status.err:", err)
		return nil, err
	}
	it, err := job.Read(ctx)
	if err != nil {
		log.Println("error iterator creation", err.Error())
		return nil, err
	}
	return it, nil
}
