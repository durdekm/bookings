package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/durdekm/bookings/pkg/config"
	"github.com/durdekm/bookings/pkg/handlers"
	"github.com/durdekm/bookings/pkg/render"
)

// If portNumber is set to ":8080" and start the program then the firewall asks
// permission to allow the binary file to accept incoming network connection.
// The message is "Möchtest du, dass das Programm main eingehende Netzwerkverbindungen akzeptiert?"
//
// Let the go program only listen on localhost:8080 (127.0.0.1:8080).
// This way the program won't need to ask for firewall traversal, and you won't
// get the message from the firewall.
// https://stackoverflow.com/questions/18978622/firewall-blocks-go-development-server
// localhost:8080 also works.
const portNumber = "127.0.0.1:8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	// don't use template cache!
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// give the render component access to this app config variable
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
