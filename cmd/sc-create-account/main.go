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
	usname := flag.String("username", "", "Email of new user.")
	hw := flag.Bool("headless", false, "Hide load indicator.")

	flag.Parse()

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    50 * time.Second,
	}

	client := &http.Client{Transport: tr}
	// wrap parameters around bson.M map under 
	// key `req`
	req := securechain.WrapReq(securechain.AuthRequest{Email : *usname})
	reqstr :=  securechain.ToString(req)

	var chn chan int
	if !*hw {
		chn = make(chan int)
		go core.DoSpin(chn)
	}
	resp, err := client.Post(securechain.Join, securechain.ContentJson, strings.NewReader(reqstr) )
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
	securechain.Log(responsedata)
	responsedata = nil
}