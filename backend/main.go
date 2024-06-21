package main

import (
	"backend/api"
	"backend/server"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.ServeIndex);
    error := http.ListenAndServe(":8080", nil);
    fmt.Println(error);
    api.GetNeets()
}
