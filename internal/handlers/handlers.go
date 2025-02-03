package handlers

import "net/http"

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("OK"))
}
