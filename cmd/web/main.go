package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fangjjcs/tooling/package/renders"

	"github.com/fangjjcs/tooling/package/config"
	"github.com/fangjjcs/tooling/package/handlers"
	"github.com/fangjjcs/tooling/package/models"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {



	// what am I going to put in the session
	gob.Register(models.Search{})

	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false // define whenever you allow to use cache or not

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplates(&app)
	const portNumber = ":8080"
	
	fmt.Printf(fmt.Sprintf("Staring application on port %s\n", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	_ = http.ListenAndServe(portNumber,nil)
	
}