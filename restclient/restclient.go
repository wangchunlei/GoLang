package main

import (
	//"bufio"
	"fmt"
	//"github.com/howeyc/gopass"
	"github.com/jmcvetta/napping"
	"log"
	//"os"
	//"strings"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {
	payload := struct {
		Logintoken struct {
			LoginType    int    `json:"LoginType"`
			TerminalType int    `json:"TerminalType"`
			Password     string `json:"Password"`
			LoginID      string `json:"LoginID"`
		} `json:"logintoken"`
	}{
		Logintoken: struct {
			LoginType    int    `json:"LoginType"`
			TerminalType int    `json:"TerminalType"`
			Password     string `json:"Password"`
			LoginID      string `json:"LoginID"`
		}{
			LoginType:    4,
			TerminalType: 4,
			Password:     "123456789a",
			LoginID:      "wangcl",
		},
	}

	res := struct {
		IsSuccess        bool   `json:"IsSuccess"`
		Message          string `json:"Message"`
		LoginToken       string `json:"LoginToken"`
		PCTocken         string `json:"PCToken"`
		DefaultAccount   string `json:"DefaultAccount"`
		IsVerifyPassword bool   `json:"IsVerifyPassword"`
	}{}

	e := struct {
		Message string
	}{}

	s := napping.Session{
	// Userinfo: url.UserPassword(username, string(passwd)),
	}
	url := "http://sd11:15391//api/BaseTerminalTerminalAPI/Login"

	resp, err := s.Post(url, &payload, &res, &e)
	checkError(err)

	println("")
	fmt.Printf("rawText: %s\n\n", resp.RawText())
	fmt.Printf("res: %s\n\n", res.LoginToken)
	println("")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
