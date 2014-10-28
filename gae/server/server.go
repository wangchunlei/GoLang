package gaeserver

import (
	"appengine"
	"appengine/urlfetch"
	"io/ioutil"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/fetch", fetch)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get("https://www.google.com/")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(response)
}

func fetch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("fetch"))
}
