package apiZeroOne

import (
	"net/http"
	"dbio"
	"fmt"
	"text/template"
	"github.com/gorilla/sessions"
	"utilities"
)

var store = sessions.NewCookieStore([]byte("chunkybacon")) // TODO: generate a secret randomly

func MemorySubmissionHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	println("submitting a memory")
	session, err := store.Get(req, "test")
	println("Got session data")


	if err != nil {
		println("There was an issue getting session data")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		panic(err)
		return
	} else {
		println("session data got successfully")
	}

	memory := utilities.ExtractMemory(req.Form)

	username, _ := session.Values["username"].(string)
	println(username)
	if err = dbio.SubmitMemory(memory, username); err != nil {
		res.Write([]byte("Successfully record your memory"))
	} else {
		res.Write([]byte("Failed to record your memory"))
		panic(err)
	}

}

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

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form.Get("username")
	password := req.Form.Get("password")
	password_confirm := req.Form.Get("password_confirm")
	if password != password_confirm {
		res.Write([]byte("Password unconfirmed! Sorry!"))
	} else {
		success := dbio.RegisterNewUser(username, password)
		if success {
			res.Write([]byte("Successfully registered " + username))
		} else {
			res.Write([]byte("Could not register username! Sorry."))
		}
	}
}



/*

        SOME HELPER FUNCTIONS

 */

func getPagePath(page string) string {
	return "pages/" + page
}
