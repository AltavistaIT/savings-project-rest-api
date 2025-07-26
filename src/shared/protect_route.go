package shared

import "net/http"

func Protect(middleware func(http.Handler) http.Handler, handlerFunc http.HandlerFunc) http.Handler {
	return middleware(handlerFunc)
}
