package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var todos = []string{}

func main() {
	http.HandleFunc("/todos", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "POST":
			// tambahkan todo baru
			todoBaru := request.URL.Query().Get("todo")
			todos = append(todos, todoBaru)
			fmt.Fprintf(writer, "Todo baru telah ditambahkan, %s\n", todoBaru)
		case "GET":
			// tampilkan semua todo
			json.NewEncoder(writer).Encode(todos)

		default:
			http.Error(writer, "Metode Request tidak diizinkan", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":1842", nil)
}
