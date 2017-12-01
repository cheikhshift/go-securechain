package securechain

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

//API base path
const prefx string = "https://sc.gophersauce.com/momentum/funcs?name=%s"

// Path to Login endpoint
var Login string = Endpoint("Login")

// Path to LoginVPhone : Login verify phone
// Use this endpoint to confirm an SMS code
// and retrieve an access token.
var VerifyLogin string = Endpoint("LoginVPhone")

// Path to Join endpoint.
var Join string = Endpoint("Join")

//Path to reset password endpoint
var ResetPassword = Endpoint("ForgotPassword") 

//Path to Encrypt endpoint.
var Encrypt = Endpoint("Encrypt") 

//Path to Decrypt endpoint.
var Decrypt = Endpoint("Decrypt") 

//Path to delete a user token.
var DeleteToken = Endpoint("DeleteToken")



type AuthRequest struct {
	Email, Password,Data string
}


// Convert specified interface to
// a json string.
// v - interface to convert to json
// return json string of interface.
func ToString(v interface{}) string {
	data, _ := json.Marshal(&v)
	return string(data)
}


// Convert json string to bson.M
func ToMap(s string) (m bson.M,err error) {
	b := []byte(s)
	err = json.Unmarshal(b, &m)
	b = nil
	return
}

// Log secure chain API response
// b - Bytes of API response.
// return bson.M of response.
// map.
func Log(b []byte) bson.M {
	var m bson.M
	err := json.Unmarshal(b, &m)
	if err != nil {
		color.Red(fmt.Sprintf("Error: %s", err.Error()))
	}

	if _, iserror := m["error"]; iserror {
		color.Red(fmt.Sprintf("Error: %s", m["error"].(string)))
		return m
	} else if _, isresp := m["resp"]; isresp {
		response := m["resp"].(map[string]interface{})
		if response["Code"].(float64) == 500 {
			color.Red(fmt.Sprintf("Error: %s", response["Message"].(string)))
		} else if response["Code"].(float64) == 500 {
			color.Yellow(fmt.Sprintf("Action: %s", response["Message"].(string)))
		} else {
			color.Green(fmt.Sprintf("Success: %s", response["Message"].(string)))
		}
		return response
	}

	return m
}


// Return bytes of *http.Response body.
// r - response to read body form.
func ReadBody(r *http.Response) (body []byte) {
	body, _ = ioutil.ReadAll(r.Body)
	return
}

// Generate a string of URL path
// to specified SecureChain API endpoint.
func Endpoint(s string) (r string) {
	r = fmt.Sprintf(prefx, s)
	return
}

// Wrap the specified interface
// arround a bson.M map, under key
// req.
func WrapReq(v interface{}) bson.M {
	nbson := bson.M{"req" : v}
	return nbson
}
