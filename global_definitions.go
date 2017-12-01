package securechain

import (
	"io/ioutil"
	"github.com/cheikhshift/form"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"runtime"
	"fmt"
)

const ContentJson string = "application/json"



//Path of folder to save session files and 
// load them.
const SessionStorePath string = "."

//Key used to encrypt session information
var Key []byte = []byte("a very very very very secret key")

// Save and encrypt the specified interface
// with the specified name. The function will
// add prefix SessionStorePath to the name specifed,
// as well as your OS's path separator.
func Save(name *string, v interface{}) error{

	str :=  form.Encrypt(Key, ToString(v))
	separator := "/"
	if strings.Contains(runtime.GOOS, "indows"){
		separator = "\\"
	}
	pathoffile := fmt.Sprintf("%s%s%s", SessionStorePath,separator, *name)
	strbytes := []byte(str)
	err := ioutil.WriteFile(pathoffile, strbytes, 0700)
	strbytes = nil
	return err
}

// Load a previously saved interface 
// with function Save(...).
func Load(name *string) (m bson.M,err error) {
	separator := "/"
	if strings.Contains(runtime.GOOS, "indows"){
		separator = "\\"
	}
	pathoffile := fmt.Sprintf("%s%s%s", SessionStorePath,separator, *name)
	data, err := ioutil.ReadFile(pathoffile)
	if err != nil {
		return
	}
	strdata := string(data)
	decryptedString := form.Decrypt(Key, strdata)
	data = nil
	m,err = ToMap(decryptedString)
	return
}

