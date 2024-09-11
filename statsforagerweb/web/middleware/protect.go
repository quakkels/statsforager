package middleware

import (
	"fmt"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func Protect(sessionManager *scs.SessionManager, next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		accountCode := sessionManager.Get(r.Context(), "accountCode")
		if accountCode == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next(w, r)
	}
}
