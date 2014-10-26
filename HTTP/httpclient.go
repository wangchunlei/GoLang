/* ClientGet
 */

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// os.Setenv("HTTP_PROXY", "http://203.233.63.168")
	// os.Setenv("HTTPS_PROXY", "http://203.233.63.168")
	url, err := url.Parse("http://life-force.appspot.com")
	checkError(err)

	proxyUrl, _ := url.Parse("http://203.233.63.168")
	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	// transport := &http.Transport{}
	client := &http.Client{Transport: transport}
	fmt.Println(url.String())
	request, err := http.NewRequest("GET", url.String(), nil)
	// only accept UTF-8
	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	checkError(err)

	response, err := client.Do(request)
	fmt.Println(response)
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Printf("got charset %s\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		os.Exit(4)
	}

	var buf [512]byte
	reader := response.Body
	fmt.Println("got body")
	for {
		n, _ := reader.Read(buf[0:])
		if n == 0 {
			break
		}
		fmt.Println(string(buf[0:n]))
	}

	os.Exit(0)
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
