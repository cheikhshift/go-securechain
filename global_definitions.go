package securechain

import (
	"io/ioutil"
	"github.com/cheikhshift/form"
	"gopkg.in/mgo.v2/bson"
	"path/filepath"
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
	pathoffile := filepath.Join(SessionStorePath, *name)
	strbytes := []byte(str)
	err := ioutil.WriteFile(pathoffile, strbytes, 0700)
	strbytes = nil
	return err
}

// Load a previously saved interface 
// with function Save(...).
func Load(name *string)  (bson.M, error) {
	var m bson.M
	pathoffile := filepath.Join(SessionStorePath, *name)

	data, err := ioutil.ReadFile(pathoffile)
	if err != nil {
		return m,err
	}
	strdata := string(data)
	decryptedString := form.Decrypt(Key, strdata)
	data = nil
	m,err = ToMap(decryptedString)
	return m,err
}

