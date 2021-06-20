package search

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fangjjcs/tooling/package/models"
)

// Response
type Response struct{
	STATUS    bool 
	MESSAGE   string 
	TYPE      string
}

func GetFormData(r *http.Request) models.Search {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	search := models.Search{
		Site: r.Form.Get("site"),
		Tool: r.Form.Get("tool"),
		Status: r.Form.Get("status"),
		Date: r.Form.Get("date"),
		Time: r.Form.Get("time"),
	}
	return search
}

// Connection with DB and Get response

var DB *sql.DB
func ConnectDB(search *models.Search) (Response, error) {
	
	log.Println(search)

	//********** DB connection ***********//

	// DB, err := sql.Open("","")
	// if err!=nil{
	// 	log.Fatal("Connection Failed")
	// 	return Response{}, err
	// }
	
	//defer DB.Close()
	//************************************//

	res, err := Post(DB,search,10)
	if err != nil{
		log.Fatal("Search Failed")
		return Response{}, err
	}

	

	return res, nil


	
}

func Post(DB *sql.DB, data *models.Search, limit int) ( Response, error){

	// query statement
	// rows, err = db.Query(....)

	// response example
	resp := Response{
		STATUS : true,
		MESSAGE : "Available!",
		TYPE : "InProcess",
	}
	return resp, nil

}