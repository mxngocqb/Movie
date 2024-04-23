package main

import "net/http"

func (app *application) enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Đặt header cho Access-Control-Allow-Origin, chỉ cho phép truy cập từ origin là "http://localhost:3000"
		w.Header().Set("Access-Control-Allow-Origin", "https://192.168.88.130:3000")

		if r.Method == "GET" || r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH"{
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Accept,Content-Type,X-CSRF-Token,Authorization,Set-Cookie")
		}

		if r.Method == "OPTIONS" {
			// Đặt các header liên quan đến CORS cho OPTIONS request
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,X-CSRF-Token,Authorization,Set-Cookie")
			return
		} else {
			// Nếu không phải là OPTIONS request, chuyển tiếp yêu cầu cho handler gốc
			h.ServeHTTP(w, r)
		}

	})
}

func (app *application) authRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, err := app.auth.GetTokenFromHeaderAndVerify(w, r)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
