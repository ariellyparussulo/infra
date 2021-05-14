package commons

import "encoding/json"

type Response struct {
	StatusCode int
	Message    string
}

var RESPONSE_STATUS = map[string]int{
	"SUCCESS":               200,
	"BAD_REQUEST":           400,
	"UNAUTHORIZED":          401,
	"FORBIDDEN":             403,
	"NOT_FOUND":             403,
	"INTERNAL_SERVER_ERROR": 500,
}

func BuildResponse(StatusCode int, Message string) []byte {
	message := Response{StatusCode, Message}
	httpResponse, err := json.Marshal(message)

	if err != nil {
		message := Response{RESPONSE_STATUS["INTERNAL_SERVER_ERROR"], "Internal Server Error"}
		httpResponse, _ := json.Marshal(message)
		return httpResponse
	}

	return httpResponse
}
