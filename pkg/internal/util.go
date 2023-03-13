package internal

import (
	"encoding/json"
)

type Response struct {
	Status string        `json:"status"`
	Data   []interface{} `json:"data"`
}

func SuccessResponse(T interface{}) []byte {
	var source []interface{}

	source = append(source, T)

	response := Response{
		Status: "Ok",
		Data:   source,
	}

	responseBytes, _ := json.Marshal(response)
	return responseBytes
}
