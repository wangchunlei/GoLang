package gaeserver

import (
	"appengine"
	"appengine/urlfetch"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestInfo struct {
	Url    string
	Header *http.Header
	Method string
	Body   *[]byte
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/fetch", fetch)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get("http://www.baidu.com/")
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
	c := appengine.NewContext(r)
	if r.Method == "GET" {
		w.Write([]byte("fetch"))
	} else {
		decoder := json.NewDecoder(r.Body)
		var t RequestInfo
		err := decoder.Decode(&t)
		if err != nil {
			c.Errorf("err: %v", err)
			return
		}
		request, err := http.NewRequest(t.Method, t.Url, bytes.NewBuffer(*t.Body))
		if err != nil {
			c.Errorf("err: %v", err)
			return
		}

		client := urlfetch.Client(c)
		resp, err := client.Do(request)
		if err != nil {
			c.Errorf("err: %v", err)
			return
		}
		defer resp.Body.Close()
		io.Copy(w, resp.Body)
	}
}
