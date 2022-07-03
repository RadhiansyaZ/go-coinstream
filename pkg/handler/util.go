package handler

import "net/http"

func SetupCorsResponse(address string, w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", address)
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
