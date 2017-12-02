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
	"io/ioutil"
	"gopkg.in/mgo.v2/bson"
	 b64 "encoding/base64"
)


func main() {
	
	text := flag.String("text", "", "Text to encrypt.")
	hw := flag.Bool("headless", false, "Hide load indicator.")
	usefile := flag.Bool("file", false, "Add flag to load encrypted data from file.")
	filename := flag.String("filename","", "Path of file with encrypted data.")
	wheredispl := flag.Bool("write-to-file", false, "Add flag to write decrypted data to file.")
	output :=  flag.String("output","", "Path of file to write decrypted data to. Leaving this flag empty will overwrite the specified file.")
	sessionfile := flag.String("session", "main.session", "Session name to use. Leave blank to use default.")

	flag.Parse()

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    50 * time.Second,
	}

	client := &http.Client{Transport: tr}
	// wrap parameters around bson.M map under 
	// key `req`
	var strdata string

	if *usefile {
		filedata,err := ioutil.ReadFile(*filename)
		if err != nil {
			color.Red(fmt.Sprintf("Error: %s", err.Error()))
			return
		}
		strdata = string(filedata)
	} else {
		strdata = *text
	}
	token,  err := securechain.Load(sessionfile)
	req := bson.M{"req": securechain.AuthRequest{ Data: strdata },"token" : token }
	reqstr :=  securechain.ToString(req)
	
	var chn chan int
	if !*hw {
		chn = make(chan int)
		go core.DoSpin(chn)
	}
	resp, err := client.Post(securechain.Decrypt, securechain.ContentJson, strings.NewReader(reqstr) )
	if !*hw {
	chn <- 1
	close(chn)
	chn = nil
	}

	if err != nil  {
		color.Red(fmt.Sprintf("Error: %s", err.Error()))
		return
	} else if resp.StatusCode == 500 {
		// this is to check HTTP status 500
		responsedata := securechain.ReadBody(resp)
		securechain.Log(responsedata)
		responsedata = nil
		return
	}

	responsedata := securechain.ReadBody(resp)

	m, err := securechain.ToMap(string(responsedata))
	if err != nil {
		color.Red(fmt.Sprintf("Error: %s", err.Error()))
		return
	}
	sm := m["resp"].(map[string]interface{})

	if sm["Code"].(float64) == 500 {
		color.Red(fmt.Sprintf("Error: %s", sm["Message"].(string)))
	} else if sm["Code"].(float64) == 200 {
		
		message, _ := b64.StdEncoding.DecodeString(sm["Message"].(string))
		if *wheredispl {
			saveto := *output
			if saveto == "" {
				saveto = *filename
			}
			err = ioutil.WriteFile(saveto, message, 0700)
			if err != nil {
				color.Red(fmt.Sprintf("Error: %s", err.Error()))
				return
			}
		} else {
			fmt.Println(string(message))
		}
	} 

	responsedata = nil
}