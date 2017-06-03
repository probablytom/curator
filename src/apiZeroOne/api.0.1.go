package apiZeroOne

import (
	"net/http"
	"dbio"
	"github.com/gorilla/sessions"
	"utilities"
	"log"
	"io/ioutil"
	"encoding/json"
	"types"
)

var store = sessions.NewCookieStore([]byte("chunkybacon")) // TODO: generate a secret randomly

func MemorySubmissionHandler(res http.ResponseWriter, req *http.Request) {
	session, err := store.Get(req, "test")
	if err != nil {
		println("There was an issue getting session data")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		panic(err)
		return
	}

	json, err := ioutil.ReadAll(req.Body)
	if err == nil {

		memory, err := utilities.ExtractJSONMemory(json)
		if err != nil {
			log.Panic("Couldn't parse JSON!")
			http.Error(res, err.Error(), http.StatusBadRequest)
		}

		username, _ := session.Values["username"].(string)
		if err = dbio.SubmitMemory(memory, username); err != nil {
			res.Write([]byte("1"))
		} else {
			res.Write([]byte("-1"))
			panic(err)
		}


	} else {
		log.Panic(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

// TODO: move to JSON request instead of forms
func LoginHandler(res http.ResponseWriter, req *http.Request) {

	session, err := store.Get(req, "test")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	dec := json.NewDecoder(req.Body)
	var details types.LoginDetails
	err = dec.Decode(details)

	if err == nil {

		username, password := details.Username, details.Password
		if dbio.CheckValidLoginCredentials(username, password) {
			session.Values["username"] = username
			session.Values["password"] = password
			session.Save(req, res)
			res.Write([]byte("Logged in"))
		} else {
			res.Write([]byte("Not logged in"))
		}

	} else {
		log.Panic(err)
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

}

// TODO: confirm password with copy?
func RegisterHandler(res http.ResponseWriter, req *http.Request) {

	dec := json.NewDecoder(req.Body)
	var details types.LoginDetails
	err := dec.Decode(details)

	if err == nil {

		username, password := details.Username, details.Password

		success := dbio.RegisterNewUser(username, password)
		if success {
			res.Write([]byte("Successfully registered " + username))
		} else {
			res.Write([]byte("Could not register username! Sorry."))
		}

	} else {
		log.Panic(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}



/*

        SOME HELPER FUNCTIONS

 */

func getPagePath(page string) string {
	return "pages/" + page
}
