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

// URL of Login endpoint.
var Login string = Endpoint("Login")

// URL of LoginVPhone endpoint: Login verify phone
// Use this endpoint to confirm an SMS code
// and retrieve an access token.
var VerifyLogin string = Endpoint("LoginVPhone")

// URL of Join endpoint.
var Join string = Endpoint("Join")

// URL of reset password endpoint.
var ResetPassword = Endpoint("ForgotPassword") 

//URL of Encrypt endpoint.
var Encrypt = Endpoint("Encrypt") 

// URL of Decrypt endpoint.
var Decrypt = Endpoint("Decrypt") 

// URL of DeleteToken endpoint.
var DeleteToken = Endpoint("DeleteToken")



type AuthRequest struct {
	Email, Password,Data string
}


// Convert specified interface to
// a json string.
// v - interface to convert to json.
// return json string of interface.
// @test
// @case bson.M{} @equal "{}"
func ToString(v interface{}) string {
	data, _ := json.Marshal(&v)
	return string(data)
}


// Convert json string to bson.M.
// @import "gopkg.in/mgo.v2/bson"
// @test
// @case "{}" @equal bson.M{}, nil
func ToMap(s string) (bson.M, error) {
	var m bson.M
	b := []byte(s)
	err := json.Unmarshal(b, &m)
	b = nil
	return m,err
}

// Log secure chain API response.
// b - Bytes of API response.
// return bson.M of response.
// map.
// @test
// @case []byte("{}") @equal bson.M{}
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
// @test
// @import "net/http"
// @import "strings"
// @import "io/ioutil"
// @case &http.Response{Body : ioutil.NopCloser(strings.NewReader("test") )} @equal []byte("test")
func ReadBody(r *http.Response)  []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

// Generate a string of URL path
// to specified SecureChain API endpoint.
// @test
// @case "Login" @equal "https://sc.gophersauce.com/momentum/funcs?name=Login"
func Endpoint(s string) string {
	r := fmt.Sprintf(prefx, s)
	return r
}

// Wrap the specified interface
// arround a bson.M map, under key
// req.
// @test
// @case bson.M{} @equal bson.M{"req":bson.M{}} 
func WrapReq(v interface{}) bson.M {
	nbson := bson.M{"req" : v}
	return nbson
}
