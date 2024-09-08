package web

import (
	"fmt"
	"net/http"
	"statsforagerweb/domain"
	"time"

	"github.com/alexedwards/scs/v2"
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

func postLoginHandler(accountsManager domain.AccountsManager, sessionManager *scs.SessionManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		model := loginModel{
			Email: r.Form.Get("email"),
		}

		otp, err := domain.NewOtpToken(10 * time.Minute)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
		sessionManager.Put(r.Context(), "LoginOtp", otp)

		validation, err := accountsManager.SendLoginMail(r.Context(), model.Email, otp.Otp)
		if err != nil {
			fmt.Println(err)
		}

		model.IsPostSuccess = validation.IsSuccess
		model.Errors = validation.ToMessagesSlice()

		if err := tplGlob.ExecuteTemplate(w, "login.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}

func getLoginConfirmHandler(
	_ domain.AccountsManager,
	sessionManager *scs.SessionManager,
) func(http.ResponseWriter, *http.Request) {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		defer sessionManager.Remove(r.Context(), "LoginOtp")
		suggestedOtp := r.PathValue("otp")
		loginOtp := sessionManager.Get(r.Context(), "LoginOtp").(domain.OtpToken)
		if loginOtp.IsValid(suggestedOtp) {
			sessionManager.RenewToken(r.Context()) // prevent session fixation
			// save authenticated user claims
			// redirect to dashboard
		}
		model := loginModel{
			IsPostSuccess: false,
			Errors:        []string{"Login unsuccessful. Register your email addres, or try again."},
		}

		if err := tplGlob.ExecuteTemplate(w, "login.html", model); err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), 500)
		}
	}
}
