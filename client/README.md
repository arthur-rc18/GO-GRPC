# Client Microservice

The client microservice is a REST API built with Gin Framework 
that provides the data of temperature and rain (precipitation) - 
minimum, maximum and instant for both - and their creation and data timestamp

## Endpoints
- There's only one endpoint in this application/microservice that is:
```http request
 GET 'domain'/api/:blockId
```
### Response Codes
 - 200 Ok - The information block exists
 - 404 NotFound - The information block doesn't exists

### Execution examples
- ##### A block that exists (B233):
```http request
 GET 'domain'/api/B233
```
Returns:
```
 200 Ok
```
```json
{
  "data_timestamp": "2022-06-28T23:05:24.799291Z",
  "created_timestamp": "2022-06-28T23:05:24.799291Z",
  "temperature_inst": "16.158027",
  "temperature_min": "14.476000",
  "temperature_max": "19.244000",
  "precipitation_inst": "0.000000"
}
```

- ##### A block that not exists:
```http request
 GET 'domain'/api/random_id
```
Returns:
```
 404 NotFound
```
```json
```