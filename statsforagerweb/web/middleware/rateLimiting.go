package middleware

import (
	"fmt"
	"net/http"

	"github.com/didip/tollbooth/v7"
	"github.com/didip/tollbooth/v7/limiter"
)

type RateLimitingMiddleware struct {
	limiter *limiter.Limiter
}

func NewRateLimitingMiddleware(limiter *limiter.Limiter) RateLimitingMiddleware {
	return RateLimitingMiddleware{limiter}
}

func (self *RateLimitingMiddleware) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in rate limiter")

		httpErr := tollbooth.LimitByRequest(self.limiter, w, r)
		if httpErr != nil {
			http.Error(w, httpErr.Message, http.StatusBadRequest)
			return
		}
		
		next.ServeHTTP(w, r)

		fmt.Println("leaving rate limiter")
	})
}
