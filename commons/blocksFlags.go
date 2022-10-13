package commons

import "flag"

var (
	PORTService = flag.String("port blocks service", "localhost:8080", "Blocks service server port")

	PORTClient = flag.String("port blocks client", "localhost:8080", "Blocks client server port")

	PORTGinAPI = flag.String("port gin api", ":8000", "Gin APi server port")

	ROUTEApi = flag.String("route api", "/api/", "api route for block")
)
