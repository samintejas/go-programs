package middleware

import (
	"context"
	"net/http"
)

const requestIdHeader = "request-id"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := "APP::IG::12345"

		if r.Header.Get(requestIdHeader) != "" {
			requestID = r.Header.Get(requestIdHeader)
		}
		ctx := context.WithValue(r.Context(), requestIdHeader, requestID)
		r = r.WithContext(ctx)

		w.Header().Set(requestIdHeader, requestID)

		next.ServeHTTP(w, r)
	})
}

func GetRequestIDFromContext(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}
	return ""
}
