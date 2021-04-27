package controller

import (
	"net/http"
	view "../view"
	"encoding/json"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			data := view.Response {
				Code : 200,
				Body : "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}