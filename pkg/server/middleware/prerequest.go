package middleware

import (
	"context"
	"encoding/json"
	"github.com/airstrik/gobase/pkg/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func PreRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestId :=middleware.GetReqID(ctx)
		log.Print("Request Handling for the Request Id ", requestId)
		_CSRFToken := r.Header.Get("X-")
		_AccountId := chi.URLParam(r, "AccountId")
		if  _CSRFToken != "" {
			// Here Comes the Account Validation and Authorization
		}
		ctx = context.WithValue(ctx, "AccountCtx", config.NewContext(_AccountId))

		// Set the Default Headers for the Response Header
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "GET" {
			var requestBody interface{}
			body := json.NewDecoder(r.Body)
			body.Decode(&requestBody)
			ctx = context.WithValue(ctx, "Body", requestBody)
		}
		w.Header().Add(middleware.RequestIDHeader, requestId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}


func processRequest(w http.ResponseWriter, r *http.Request) {

}