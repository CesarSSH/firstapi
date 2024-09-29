package routes

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Escribe el texto "Hello World" como respuesta
	w.Write([]byte("My server is running"))
}
