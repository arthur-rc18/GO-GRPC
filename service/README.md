# Service Module

Contains all database connection and is responsible for sending the query searching for the specific blockId.

## Package/structure summary

. database: 
    create a new BigQuery client
    sets up the connection to database
    
. model:
    creates the query to be sended in bigQuery
    consults the pixel table

. main:
    runs the service
    implements protobuf interface    