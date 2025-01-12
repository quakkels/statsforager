package web

import (
	"fmt"
	"net/http"
	"statsforagerweb/domain"
)

func getAppHandler(sitesManager domain.SitesManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// override for development
		currentAccountKey := "me@example.com"

		sites, err := sitesManager.GetSites(r.Context(), currentAccountKey)
		if err != nil {
			fmt.Println("sites.go: getAppHandler>", err.Error())
			http.Error(w, "This should never happen, but it has", 500)
			return
		}

		model := struct {
			Form  domain.Site
			Sites []domain.Site
		}{
			Sites: sites,
		}

		render(w, r, "sites.html", model)
	}
}

func postAppManageSiteSave(sitesManager domain.SitesManager) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get form data
		r.ParseForm()

		model := domain.Site {
			OwnerAccount: "me@example.com",
			Domain: r.Form.Get("siteDomain"),
			SiteName: r.Form.Get("siteName"),
			SiteKey: "this is a thing",
		}

		// validate form data
		validationErrors := []string{}
		if model.Domain == "" {
			validationErrors= append(validationErrors, "Missing domain")
		}

		if len(validationErrors) > 0 {
			renderRawBadRequest(w, "_errors.html", validationErrors)
			return
		}

		// save site
		// todo save

		// return site html
		renderRaw(w, "_siteRow.html", model)
	}
}

/*
func getSiteSaveHandler(sitesManager domain.SitesManager) func(http.ResponseWriter, *http.Request) {
return func(w http.ResponseWriter, r *http.Request) {

currentSiteKey := r.URL.Query().Get("SiteKey")
currentDomain := r.Form.Get("Domain")

accountCode, ok := GetAccountCode(r.Context())
if !ok {
fmt.Println("sites.go: getSiteSaveHandler> ERROR: the account code was not found.")
http.Error(w, "This should never happen, but it has", 500)
return
}


site := domain.Site{
SiteKey: uuid.New().String(),
SiteOwner: accountCode,
Domain: currentDomain,
}

if err != nil {
fmt.Println("sites.go: getSiteSaveHandler> couldn't get site:", err.Error())
http.Error(w, "This should never happen, but it has", 500)
return
}


}
}
*/
