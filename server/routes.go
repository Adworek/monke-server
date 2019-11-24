package server

import (
    "net/http"
    "strings"
)

func RouteUserTree(response http.ResponseWriter, request *http.Request) {
    switch request.Method {
    case "GET":
        getRouteUserTree(response, request)
        return
    default:
        // handle http errors util
    }
}

func getRouteUserTree(response http.ResponseWriter, request *http.Request) {
    var path_items []string = strings.Split(request.URL.Path, "/")

    if path_items[2] == "by_nick" && len(path_items) == 4 {
        getUserByNick(path_items[3], response, request)
        return
    }
}
