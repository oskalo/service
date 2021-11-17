package middlewares

import (
	"fmt"
	"net/http"
)

func LogEndPointCalls(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logText := fmt.Sprintf("RemoteAddr: %s, method: %s, RequestURI: %s",r.RemoteAddr, r.Method, r.RequestURI)
		fmt.Println(logText)
		nextHandler.ServeHTTP(w, r)
	})
}
