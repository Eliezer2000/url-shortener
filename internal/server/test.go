package server

import "net/http"

func TestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("World"))
	}
}