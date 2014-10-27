/* ClientGet
 */

package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	url, err := url.Parse("https://" + os.Args[1])
	checkError(err)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	fmt.Println(url.String())
	request, err := http.NewRequest("GET", url.String(), nil)
	checkError(err)
	request.Host = "life-force.appspot.com"
	fmt.Println(request.Host)
	response, err := client.Do(request)
	checkError(err)
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
