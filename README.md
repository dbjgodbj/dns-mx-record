# DNS MX Record Web Service Test

A web application that displays the IP addresses and mail hosts associated with a user-entered domain name.

## Requirements
· The application should be written in either Python or Go

· The application should listen on port 8080

· When visiting the root resource of your app (http://localhost:8080/), display a form asking for a domain name (e.g. dyn.com)

· When that value is submitted, display the IP addresses associated with the domain and the hosts associated with that domain's DNS MX records.

· If the results page is accessed with the HTTP Accept header set to "application/json", render a JSON response instead of HTML

    o Note, in this case, the Accept header will only be "application/json". No further content negotiation settings will be asked for.

## Getting Started

 1 - There is a contants.go file used to store the URL Port (default is 8080) and the Index template path
 2 - The program can be executed in the project root path (path where the main.go file is available) using the command below:
 ```
 go run main.go
 ```
 3 - The build process looking to generate the executable can be done as:
 ```
 go build
 go install
 ```
 The main file executable will be installed on $GOPATH/bin and the dependecies at $GOPATH/pkg

### Prerequisites

 This application is using the libraries below as dependency:

 1 - Golang logging library - github.com/op/go-logging
 ```
 $ go get github.com/op/go-logging
 ```
 2 - gorilla/mux - github.com/gorilla/mux
 ```
 $ go get -u github.com/gorilla/mux
 ```

## Authors

* **Edson Martins** - *Initial version*

## Future

* Remove characters {} from MX records when retrieving the MX data and showing on the html list at index page
* Add unit test module for json call
