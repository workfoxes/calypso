package middleware

import (
	"context"
	"github.com/airstrik/gobase/pkg/config"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

func PreRequest(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var userId string
		requestId := middleware.GetReqID(ctx)
		log.Print("Request Handling for the Request Id ", requestId)
		_CSRFToken := r.Header.Get("X-")
		if _CSRFToken != "" {
			// Here Comes the Account Validation and Authorization
		}
		config := config.LoadContext(r, userId)
		ctx = context.WithValue(ctx, "AccountCtx", config)
		// Set the Default Headers for the Response Header
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set(middleware.RequestIDHeader, requestId)
		next.ServeHTTP(w, r.WithContext(ctx))
		config.DB.DB.Commit()
	}
	return http.HandlerFunc(fn)
}
