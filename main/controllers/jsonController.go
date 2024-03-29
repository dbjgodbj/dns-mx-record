package controllers

import (
	"encoding/json"
	"github.com/op/go-logging"
	"github.com/kururu-br/dns-mx-record/main/network"
	"net/http"
	"fmt"
	"bytes"
)

type dns_json struct {
	Url string
}

var log = logging.MustGetLogger("controllers")

// - ------------------------------------------------------------------------------------------------------------------
// - Json handler used to retrieve the IP and MX data to the Json format
// - ------------------------------------------------------------------------------------------------------------------
func JsonHandler(w http.ResponseWriter, r *http.Request) {

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


		log.Debugf("JsonHandler started to research and build a json message of IP and MX data from %s domain", request.Url)

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