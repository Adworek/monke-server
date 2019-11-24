package util

import (
	"monke-server/models"
	"encoding/json"
	"fmt"
    "net/http"
)

func HTTPResponseJson(response http.ResponseWriter, response_map map[string]interface{}) {
	var parse_error error
	var response_data []byte

	response_data, parse_error = json.Marshal(response_map)

	if parse_error != nil {
		// todo handle err 500
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	fmt.Fprint(response, string(response_data))
}

func HTTPResponseUser(response http.ResponseWriter, user models.User) {
	var parse_error error
	var response_data []byte

	response_data, parse_error = json.Marshal(user)

	if parse_error != nil {
		// todo
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(200)
	fmt.Fprint(response, string(response_data))
}
