package api

import (
	"encoding/json"
	"errors"
)

// Define some known errors
var errUnauthorized = errors.New("invalid API key")

// Response represents an API response
type Response struct {
	Type string `json:"type"`
	Value string `json:"value"`
}

// errorResponse returns a marshalled error response
func errorResponse(err error) []byte {
	response := Response{
		Type: "error",
		Value: err.Error(),
	}
	value, _ := json.Marshal(response)
	return value
}

// successResponse returns a marshalled success response
func successResponse(content string) []byte {
	response := Response{
		Type: "success",
		Value: content,
	}
	value, _ := json.Marshal(response)
	return value
}
