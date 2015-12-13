package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

// Cache templates
var templates = template.Must(template.ParseFiles("home.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func main() {
	// test()
	httpHost := os.Getenv("HOST")
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8086"
	}

	InitMQTTClient()

	go h.run()
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ws", serveWs)
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	log.Printf("PetFeeder server listening on %s:%s\n", httpHost, httpPort)
	http.ListenAndServe(httpHost+":"+httpPort, nil)
}
