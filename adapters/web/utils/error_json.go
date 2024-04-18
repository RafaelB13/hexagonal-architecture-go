package utils

import "encoding/json"

func ErrorJson(message string) []byte {
	errMsg := struct {
		Message string `json:"message"`
	}{
		Message: message,
	}

	response, err := json.Marshal(errMsg)
	if err != nil {
		return []byte(err.Error())
	}
	return response
}
