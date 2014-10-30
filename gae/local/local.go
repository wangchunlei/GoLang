package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	//"net"
	"net/http"
	//"strings"
)

type RequestInfo struct {
	Url    string
	Header *http.Header
	Method string
	Body   *[]byte
}

func Post(req *http.Request) *http.Response {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	rawrequest := &RequestInfo{
		Url:    req.URL.String(),
		Header: &req.Header,
		Method: req.Method,
		Body:   &body,
	}
	rJson, _ := json.Marshal(rawrequest)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	request, err := http.NewRequest("POST", "https://64.233.169.106/fetch", bytes.NewBuffer(rJson))
	request.Host = "gaeofgo.appspot.com"
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8888", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	proxy.OnRequest().DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		if req.URL.Scheme == "https" {
			req.URL.Scheme = "http"
			req.URL.Host = req.Host + ":80"
		}
		log.Println(req.URL)
		return req, Post(req)
	})

	http.ListenAndServe(*addr, proxy)
}
