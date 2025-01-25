package docker_practice

import (
	"fmt"
	"net/http"
)

package main

import (
"fmt"
"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, это твое первое приложение на Go в Docker!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}