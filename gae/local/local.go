package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	//"net/url"
	"bytes"
	"io"
	"io/ioutil"
	"time"
)

type RequestInfo struct {
	Url    string
	Header *http.Header
	Method string
	Body   *[]byte
}

func OnRequest(w http.ResponseWriter, req *http.Request) {
	Post(w, req)
}
func Post(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	rawrequest := &RequestInfo{
		Url:    req.RequestURI,
		Header: &req.Header,
		Method: req.Method,
		Body:   &body,
	}
	rJson, _ := json.Marshal(rawrequest)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	request, err := http.NewRequest("POST", "https://0.0.0.0/fetch", bytes.NewBuffer(rJson))
	request.Host = "gaeofgo.appspot.com"
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	io.Copy(w, response.Body)
}
func main() {
	addr := flag.String("l", ":8888", "on which address should the proxy listen")
	flag.Parse()
	// target, err := url.Parse("http://192.168.70.118:8087")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	http.HandleFunc("/", OnRequest)
	s := &http.Server{
		Addr:           *addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
