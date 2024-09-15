package web

import (
	"net/http"
	"statsforagerweb/domain"
)

type registerModel struct {
	Email         string
	IsPostSuccess bool
	Errors        []string
	Token         string
}

func getRegisterHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		model := registerModel{}
		render(w, r, "register.html", model)
	}
}

func postRegisterHandler(accountsManager domain.AccountsManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		model := registerModel{
			Email: r.Form.Get("email"),
		}

		validationResult, err := accountsManager.RegisterEmail(r.Context(), model.Email)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		model.IsPostSuccess = validationResult.IsSuccess
		model.Errors = validationResult.ToMessagesSlice()
		render(w, r, "register.html", model)
	}
}
