package web

import (
	"fmt"
	"net/http"
	"statsforagerweb/domain"
)

type registerModel struct {
	Email         string
	IsPostSuccess bool
	Errors        []string
}

func getRegisterHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		model := registerModel{}

		if err := tplGlob.ExecuteTemplate(w, "register.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
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

		if err := tplGlob.ExecuteTemplate(w, "register.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}
