package webHandlers

import (
	"net/http"
	"dbio"
	"github.com/gorilla/sessions"
	"text/template"
	"utilities"
	"fmt"
)

// TODO: generate a secret randomly
var store = sessions.NewCookieStore([]byte("chunkybacon"))

// TODO: change this to use the API
func LoginHandler(res http.ResponseWriter, req *http.Request) {

	req.ParseForm()
	session, err := store.Get(req, "test")

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	username, password := req.Form.Get("username"), req.Form.Get("password")
	if dbio.CheckValidLoginCredentials(username, password) {
		session.Values["username"] = username
		session.Values["password"] = password
		session.Save(req, res)
		res.Write([]byte("Logged in"))
	} else {
		res.Write([]byte("Not logged in"))
	}

}


// TODO: change this to use the API
func AddHandler(res http.ResponseWriter, req *http.Request) {

	// We don't need to parse a form here, just throw up the submission template...
	if template_render, parse_err := template.ParseFiles(utilities.GetPagePath("add_page.html")); parse_err == nil {
		template_render.Execute(res, nil)
	} else {
		panic(parse_err)
	}
}

func ContentRetrievalHandler(res http.ResponseWriter, req *http.Request) {

	session, err := store.Get(req, "test")
	if err != nil {panic(err)}

	username, _ := session.Values["username"].(string)

	if template_render, parse_err := template.ParseFiles(utilities.GetPagePath("content_list.html")); parse_err == nil {
		template_render.Execute(res, dbio.RetrieveUserMemories(username))
	} else {
		res.Write([]byte("Could not generate page from content_list.html! Sorry."))
		fmt.Println(parse_err)
	}

}

