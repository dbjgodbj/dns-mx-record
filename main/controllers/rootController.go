package controllers

import (
	"net/http"
	"html/template"
	"github.com/kururu-br/dns-mx-record/main/constants"
)

// - ------------------------------------------------------------------------------------------------------------------
// - Root handler used to show the index html page
// - ------------------------------------------------------------------------------------------------------------------
func RootHandler(w http.ResponseWriter, r *http.Request) {

	log.Debug("RootHandler started and redirecting to the index page")

	t,_ := template.ParseFiles(constants.INDEX)
	t.Execute(w, nil)
}