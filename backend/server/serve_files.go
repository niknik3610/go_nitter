package server

import (
	"net/http"
	"os"
)

func ServeIndex(w http.ResponseWriter, req *http.Request) {
	data, err := os.ReadFile("../frontend/index.html")
	if err != nil {
		panic(err)
	}
	w.Write(data)
}
