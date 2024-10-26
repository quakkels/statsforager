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
		render(w, r, "login.html", model)
	}
}

func postLoginHandler(accountsManager domain.AccountsManager, sessionManager *scs.SessionManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		model := loginModel{
			Email: r.Form.Get("email"),
		}

		otp, err := domain.NewOtpToken(model.Email, 10*time.Minute)
		fmt.Println("login.go: postLoginHandler> otp.Otp:", otp.Otp)
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
		fmt.Println("model.IsPostSuccess:", model.IsPostSuccess)
		model.Errors = validation.ToMessagesSlice()
		render(w, r, "login.html", model)
	}
}

func getLoginConfirmHandler(sessionManager *scs.SessionManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer sessionManager.Remove(r.Context(), "LoginOtp")
		suggestedOtp := r.PathValue("otp")
		fmt.Println("login.go: getLoginConfirmHandler> suggestedOtp:", suggestedOtp)

		var loginOtp domain.OtpToken
		if tempotp, ok := sessionManager.Get(r.Context(), "LoginOtp").(domain.OtpToken); !ok {
			fmt.Println("login.go: getLoginConfirmHandler> sessiontManager: couldn't get 'LoginOtp' as domain.OtpToken")
			http.Error(w, "Invalid session or OTP", http.StatusUnauthorized)
			return
		} else {
			loginOtp = tempotp
		}
		fmt.Println("login.go: getLoginConfirmHandler> loginOtp:", loginOtp.Otp)

		if loginOtp.IsValid(suggestedOtp) {
			sessionManager.RenewToken(r.Context()) // prevent session fixation
			sessionManager.Put(r.Context(), "accountCode", loginOtp.AccountCode)
			http.Redirect(w, r, "/app/dashboard", http.StatusSeeOther)
		}
		model := loginModel{
			IsPostSuccess: false,
			Errors: []string{
				"Login unsuccessful.",
				"Make sure you're using a registered email, and you follow the login link before it expires.",
			},
		}
		render(w, r, "login.html", model)
	}
}

func getLogoutHandler(sessionManager *scs.SessionManager) func(http.ResponseWriter, *http.Request) {
	return func(
		w http.ResponseWriter,
		r *http.Request,
	) {
		sessionManager.Clear(r.Context())
		sessionManager.RenewToken(r.Context())
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
