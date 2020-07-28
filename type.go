package main

import "net/http"

type response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type handler func(w http.ResponseWriter, r *http.Request)
