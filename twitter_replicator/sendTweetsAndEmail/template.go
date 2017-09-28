package main

import (
	"encoding/json"
	"net/http"
)

func serveTemplate(res http.ResponseWriter, req *http.Request, templateName string) {
	memItem, err := getCachedUser(req)
	if err != nil {
		// not logged in
		tpl.ExecuteTemplate(res, templateName, SessionData{})
		return
	}
	// logged in
	var sd SessionData

	// store the session data
	json.Unmarshal(memItem.Value, &sd)
	sd.LoggedIn = true
	tpl.ExecuteTemplate(res, templateName, &sd)
}
