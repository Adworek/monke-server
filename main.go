package main

import (
	"monke-server/server"
	"monke-server/storage"
	"net/http"
)

func main() {
	storage.NewUser("foobar", "foobar2000")

	http.HandleFunc("/user/", server.RouteUserTree)

	http.ListenAndServe(":8000", nil)
}
