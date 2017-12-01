package main

import (
	"net/http"
	"fmt"
	"github.com/cheikhshift/securechain-cmd"
	"flag"
	"time"
	"strings"
	"github.com/fatih/color"
	"github.com/cheikhshift/gos/core"
)


func main() {
	usname := flag.String("username", "", "User's SecureChain username.")
	pw := flag.String("password", "", "User's SecureChain password")
	hw := flag.Bool("headless", false, "Hide load indicator.")
	sessionfile := flag.String("session", "main.session", "Name of session. Leave blank to use default.")

	flag.Parse()

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    50 * time.Second,
	}

	client := &http.Client{Transport: tr}
	// wrap parameters around bson.M map under 
	// key `req`
	req := securechain.WrapReq(securechain.AuthRequest{Email : *usname, Password :*pw})
	reqstr :=  securechain.ToString(req)

	
	var chn chan int
	if !*hw {
		chn = make(chan int)
		go core.DoSpin(chn)
	}
	resp, err := client.Post(securechain.Login, securechain.ContentJson, strings.NewReader(reqstr) )
	if !*hw {
	chn <- 1
	close(chn)
	chn = nil
	}

	if err != nil  {
		color.Red(fmt.Sprintf("Error: %s", err.Error()))
		return
	} else if resp.StatusCode == 500 {
		responsedata := securechain.ReadBody(resp)
		securechain.Log(responsedata)
		responsedata = nil
		return
	}

	responsedata := securechain.ReadBody(resp)
	m := securechain.Log(responsedata)
	if m["Code"].(float64) == 202 {
		fmt.Println(fmt.Sprintf("Use command `sc-confirm-phone -code=$CODE -id=%s` to send verification code.", m["Result"].(string) ) )
		fmt.Println("Replace $CODE with the verification code sent to your user's phone.")
	} else if m["Code"].(float64) == 200 {
		err := securechain.Save(sessionfile, m["Result"])
		if err != nil  {
		color.Red(fmt.Sprintf("Error: %s", err.Error()))
		return
		}
		fmt.Println("Access token saved to store. Your session's name is `", *sessionfile,"`." )
	}
	responsedata = nil
}