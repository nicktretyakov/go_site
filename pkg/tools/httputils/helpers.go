package httputils

import (
	"net/http"
	"strings"
)

// StatusOk return 200 response
func StatusOk(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// get bearer header and return token from it
func GetAuthorizationHeader(r *http.Request) (string, bool) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", false
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", false
	}

	if strings.Trim(headerParts[0], ": ") != "Bearer" {
		return "", false
	}

	return headerParts[1], true
}

// AccessControl add headers to response
func AccessControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

// OptionsHandler for handle OPTIONS method
func OptionsHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	origin := request.Header.Get("Origin")
	if origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}
	w.WriteHeader(http.StatusOK)
}
