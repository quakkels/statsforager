package web

import (
	"fmt"
	"net/http"
	"statsforagerweb/domain"
)

type loginModel struct {
	Email         string
	IsPostSuccess bool
	Errors        []string
}

func getLoginHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		model := loginModel{}

		if err := tplGlob.ExecuteTemplate(w, "login.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}

func postLoginHandler(accountsManager domain.AccountsManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("r.Form.Get(\"email\"):", r.Form.Get("email"))
		model := loginModel{
			Email: r.Form.Get("email"),
			IsPostSuccess: true,
		}

		err := accountsManager.SendLoginMail(r.Context(), model.Email)
		if err != nil {
			fmt.Println(err)
		}

		if err := tplGlob.ExecuteTemplate(w, "login.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}
