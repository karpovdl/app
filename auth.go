package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func authorized(pass handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			jsonResponse(w, http.StatusUnauthorized, "authorization failed")
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !validate(pair[0], pair[1]) {
			jsonResponse(w, http.StatusUnauthorized, "authorization failed")
			return
		}

		pass(w, r)
	})
}

func validate(username, password string) bool {
	if username == appAuthUser && password == appAuthPassword {
		return true
	}
	return false
}
