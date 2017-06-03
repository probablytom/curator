package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"text/template"
	"dbio"
	"apiZeroOne"
	"webHandlers"
	"utilities"
)

func main() {
	dbio.BeginConnection("")
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", rootHandler)

	// API ENDPOINTS
	rtr.HandleFunc("/0.1/login", apiZeroOne.LoginHandler).Methods("POST")
	rtr.HandleFunc("/0.1/register", apiZeroOne.RegisterHandler).Methods("POST")
	rtr.HandleFunc("/0.1/add", apiZeroOne.MemorySubmissionHandler).Methods("POST")

	// WEB ENDPOINTS
	rtr.HandleFunc("/login", webHandlers.LoginHandler).Methods("GET")
	rtr.HandleFunc("/0.1/add", webHandlers.AddHandler).Methods("GET")
	rtr.HandleFunc("/0.1/content", webHandlers.ContentRetrievalHandler).Methods("GET")

	http.ListenAndServe(":8080", rtr)
}

func rootHandler(res http.ResponseWriter, req *http.Request) {
	if template_render, parse_err := template.ParseFiles(utilities.GetPagePath("root.html")); parse_err == nil {
		template_render.Execute(res, nil)
	} else {
		res.Write([]byte("Could not generate page from root.html! Sorry."))
		fmt.Println(parse_err)
	}
}


/*

        SOME HELPER FUNCTIONS

 */
