package middleware

import (
	"context"
	"github.com/airstrik/gobase/pkg/config"
	"net/http"
)

func Autherization(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_CSRFToken := r.Header.Get("")
		if _CSRFToken != "" {
			// Here Comes the Account Validation and Autherization
		}
		ctx = context.WithValue(ctx, "context", config.LoadContext(r, ""))
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
