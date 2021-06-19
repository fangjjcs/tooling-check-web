package handlers

import (
	"log"
	"net/http"

	"github.com/fangjjcs/tooling/package/config"
	"github.com/fangjjcs/tooling/package/form"
	"github.com/fangjjcs/tooling/package/models"
	"github.com/fangjjcs/tooling/package/renders"
	"github.com/fangjjcs/tooling/package/search"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("home")

	// send data to the template
	renders.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
	})
}


// MakeReservation is the handler for the MakeReservation page
func (m *Repository) Search(w http.ResponseWriter, r *http.Request) {

	log.Println("search")

	post := search.GetFormData(r)
	data := make(map[string]interface{})
	data["search"] = post

	// send data to the template
	renders.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Form: form.New(nil),
		Data: data,
	})
}



// PostMakeReservation is the handler for Post the reservation form
func (m *Repository) PostSearch(w http.ResponseWriter, r *http.Request) {
	log.Println("post")
	post := search.GetFormData(r)
	form := form.New(r.PostForm)
	form.Required("site","tool","status","date","time")
	
	data := make(map[string]interface{})
	data["search"] = post

	
	if !form.Valid(){ 
		renders.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return 
	}

	// Connection with DB
	response,err := search.ConnectDB(&post)
	if err!=nil{
		log.Fatal("Connection DB failed")
		return
	}
	// Set response to template
	data["response"] = response
	renders.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Form: form,
		Data: data,
	})
	
}




