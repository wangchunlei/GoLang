package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"github.com/elazarl/goproxy"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
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
		Url:    strings.Replace(req.RequestURI, ":443", "", 1),
		Header: &req.Header,
		Method: req.Method,
		Body:   &body,
	}
	rJson, _ := json.Marshal(rawrequest)
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	request, err := http.NewRequest("POST", "https://203.233.63.168/fetch", bytes.NewBuffer(rJson))
	request.Host = "gaeofgo.appspot.com"
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	return response
}

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8888", "proxy listen address")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	proxy.OnRequest().HijackConnect(func(req *http.Request, client net.Conn, ctx *goproxy.ProxyCtx) {
		defer func() {
			if e := recover(); e != nil {
				ctx.Logf("error connecting to remote: %v", e)
				client.Write([]byte("HTTP/1.1 500 Cannot reach destination\r\n\r\n"))
			}
			client.Close()
		}()

		log.Println(req.RequestURI)
		ctx.Resp.Body = Post(req).Body
	})
	http.ListenAndServe(*addr, proxy)
}
