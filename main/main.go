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
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"
	"bytes"

	"github.com/op/go-logging"
	"github.com/kururu-br/dns-mx-record/main/network"
	"github.com/kururu-br/dns-mx-record/main/constants"
	"github.com/gorilla/mux"
)

var log = logging.MustGetLogger("main")

type dns_json struct {
	Url string
}

// - ------------------------------------------------------------------------------------------------------------------
// - Json handler used to retrieve the IP and MX data to the Json format
// - ------------------------------------------------------------------------------------------------------------------
func jsonHandler(w http.ResponseWriter, r *http.Request) {

	// - --------------------------------------------------------------------------------------------------------------
	// - Get Accept header (e.g. "application/json")
	// - --------------------------------------------------------------------------------------------------------------
	accept := r.Header.Get("Accept")

	switch accept {

	// - --------------------------------------------------------------------------------------------------------------
	// - Requirement: If the results page is accessed with the HTTP Accept header set to
	// - "application/json", render a JSON response instead of HTML
	// - --------------------------------------------------------------------------------------------------------------
	case "application/json":

		decoder := json.NewDecoder(r.Body)

		var request dns_json

		err := decoder.Decode(&request)


		log.Debugf("Called jsonHandler function and with input parameter domain as %s", request.Url)

		// Research the IP and MX data based on an URL parameter
		Dns, err := network.GetDNS(request.Url)
		if err != nil {
			log.Errorf("The DNS get function failed with error %s and will return to the web client the error %v",
				        err.Error(), http.StatusInternalServerError)

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// - ----------------------------------------------------------------------------------------------------------
		// - Handler Accept header as JSON
		// - ----------------------------------------------------------------------------------------------------------
		jsonData, _ := json.Marshal(Dns)

		// - ----------------------------------------------------------------------------------------------------------
		// - Render a JSON response instead of HTML
		// - ----------------------------------------------------------------------------------------------------------
		fmt.Fprint(w, bytes.NewBuffer(jsonData))
		return
	}
}

// - ------------------------------------------------------------------------------------------------------------------
// - Search handler used to manage the search form POST
// - ------------------------------------------------------------------------------------------------------------------
func searchHandler(w http.ResponseWriter, r *http.Request) {

	// - --------------------------------------------------------------------------------------------------------------
	// - Retrieve the HTML form parameter of POST method
	// - --------------------------------------------------------------------------------------------------------------
	url := r.FormValue("entry-domain")

	log.Debugf("Called searchHandler function and retrieving the parameter entry-domain as %s", url)

	// - --------------------------------------------------------------------------------------------------------------
	// - Get the IP and MX data based on the URL parameter
	// - --------------------------------------------------------------------------------------------------------------
	Dns, err  := network.GetDNS(url)
	if err != nil {
		log.Errorf("The DNS get function failed with error %s and will return to the web client the error %v",
			        err.Error(), http.StatusInternalServerError)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t,_ := template.ParseFiles(constants.INDEX)
	t.Execute(w, Dns)

}

// - ------------------------------------------------------------------------------------------------------------------
// - Root handler used to show the index html page
// - ------------------------------------------------------------------------------------------------------------------
func rootHandler(w http.ResponseWriter, r *http.Request) {

	log.Debug("Called rootHandler function and redirecting to the index page")

	t,_ := template.ParseFiles(constants.INDEX)
	t.Execute(w, nil)

}

// - ------------------------------------------------------------------------------------------------------------------
// - Entry point from golang application
// - ------------------------------------------------------------------------------------------------------------------
func main() {

	rtr := mux.NewRouter()
	// - --------------------------------------------------------------------------------------------------------------
	// - Router from / url and redirect to the root handler
	// - --------------------------------------------------------------------------------------------------------------
	rtr.HandleFunc("/",       rootHandler  )

	// - --------------------------------------------------------------------------------------------------------------
	// - Router of /search url and redirect to the search handler when submitting the HTML form
	// - --------------------------------------------------------------------------------------------------------------
	rtr.HandleFunc("/search", searchHandler)

	// - --------------------------------------------------------------------------------------------------------------
	// - Router of /json url and redirect to the json handler used to retrieve the IP and MX data in JSON format
	// - --------------------------------------------------------------------------------------------------------------
	rtr.HandleFunc("/json",   jsonHandler).Methods("POST")

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
	fmt.Println(http.ListenAndServe(constants.PORT, nil))
}