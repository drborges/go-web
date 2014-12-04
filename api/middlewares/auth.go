package middlewares

import "net/http"

func Auth(req *http.Request, res http.ResponseWriter) {
	if req.Header.Get("X-Auth-Token") != "secret123" {
		res.WriteHeader(http.StatusUnauthorized)
	}
}
