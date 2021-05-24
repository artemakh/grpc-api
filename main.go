package main

import (
	"net/http"
	"uid/client"
)

func main() {
	http.HandleFunc("/get_method/", client.GetUsersId)

	http.ListenAndServe("127.0.0.1:8000", nil)
}
