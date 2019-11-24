package server

import (
	"monke-server/models"
	"monke-server/storage"
	"monke-server/util"
	"net/http"
)

func getUserByNick(nick string, response http.ResponseWriter, request *http.Request) {
	var user models.User
	var exists bool
	user, exists = storage.UserByNick(nick)

	if !exists {
		// handle 404
		return
	}

	util.HTTPResponseUser(response, user)
    return
}
