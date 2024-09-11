package middleware

import (
	"context"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type HydrateAccountMiddleware struct {
	sessionManager *scs.SessionManager
}

func NewHydrateAccountMiddleware(sessionManager *scs.SessionManager) HydrateAccountMiddleware {
	return HydrateAccountMiddleware{sessionManager: sessionManager}
}

func (self *HydrateAccountMiddleware) Apply(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accountCode := self.sessionManager.Get(r.Context(), "accountCode")
		if accountCode != nil {
			r = r.WithContext(context.WithValue(r.Context(), "accountCode", accountCode))
		}

		next.ServeHTTP(w, r)
	})
}
