/* - ------------------------------------------------------------------------------------------------------------------
   - @author: Edson Martins
   - @version: 0.1
   - @date: May 02, 2018
   -
   - Module: main.go
   -
   - Description: This is a web application that displays the IP addresses and mail hosts associated with an
   -              user-entered domain name.
   -
   - Requirements:
   - . The application should be written in either Python or Go.
   - . The application should listen on port 8080
   - . When visiting the root resource of your app (http://localhost:8080/),
   -   display a form asking for a domain name (e.g. dyn.com)
   - . When that value is submitted, display the IP addresses associated with the domain and the hosts
   -   associated with that domain's DNS MX records.
   - . If the results page is accessed with the HTTP Accept header set to "application/json",
   -   render a JSON response instead of HTML
   - ...Note, in this case, the Accept header will only be "application/json".
   -
   - -----------------------------------------------------------------------------------------------------------------*/
package main

import (
	"net/http"

	"github.com/op/go-logging"
	"github.com/kururu-br/dns-mx-record/main/constants"
	"github.com/gorilla/mux"
	"github.com/kururu-br/dns-mx-record/main/controllers"
)

var log = logging.MustGetLogger("main")

// - ------------------------------------------------------------------------------------------------------------------
// - Entry point from golang application
// - ------------------------------------------------------------------------------------------------------------------
func main() {

	log.Infof("Started DNS MX Record Application. URL Port [%v] ", constants.PORT)

	rtr := mux.NewRouter()
	// - --------------------------------------------------------------------------------------------------------------
	// - Router from / url and redirect to the root handler
	// - --------------------------------------------------------------------------------------------------------------
	rtr.HandleFunc("/",       controllers.RootHandler)

	// - --------------------------------------------------------------------------------------------------------------
	// - Router of /search url and redirect to the search handler when submitting the HTML form
	// - --------------------------------------------------------------------------------------------------------------
	rtr.HandleFunc("/search", controllers.SearchHandler)

	// - --------------------------------------------------------------------------------------------------------------
	// - Router of /json url and redirect to the json handler used to retrieve the IP and MX data in JSON format
	// - --------------------------------------------------------------------------------------------------------------
	rtr.HandleFunc("/json",   controllers.JsonHandler).Methods("POST")

	// - --------------------------------------------------------------------------------------------------------------
	// - Define the web root path for css, js and any asset used on the HTML templates
	// - --------------------------------------------------------------------------------------------------------------
		rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// - --------------------------------------------------------------------------------------------------------------
	// - Start HTTP handle
	// - --------------------------------------------------------------------------------------------------------------
	http.Handle("/",rtr)

	// - --------------------------------------------------------------------------------------------------------------
	// Application will listen port <nnnn>, where nnnn is configured on constants package
	// - --------------------------------------------------------------------------------------------------------------
	log.Fatal(http.ListenAndServe(constants.PORT, nil))
}
