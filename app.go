package main

import (
	"log"
	"net/http"
	"os"
	"html/template"
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"
)

const (
	DEFAULT_PORT = "8080"
)

var index = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/index.html",
))

var converter = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/us2sane.html",
))

var messagegen = template.Must(template.ParseFiles(
  "templates/_base.html",
  "templates/motd.html",
))
func helloworld(w http.ResponseWriter, req *http.Request) {
  index.Execute(w, nil)
}

func us2sane(w http.ResponseWriter, req *http.Request) {
  converter.Execute(w, nil)
}

func motd(w http.ResponseWriter, req *http.Request) {
  messagegen.Execute(w, nil)
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", helloworld)
	http.HandleFunc("/converter", us2sane)
	http.HandleFunc("/motd", motd)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)
}
