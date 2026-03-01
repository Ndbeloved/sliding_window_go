package router

import (
	"net/http"

	"github.com/Ndbeloved/rate-limiter-go/internals/handler"
)

func NewRouter(rlMiddleware func(http.Handler) http.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/health", rlMiddleware(http.HandlerFunc(handler.Health)))
	return mux
}
