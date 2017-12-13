package securechain

import (
	"io/ioutil"
	"github.com/cheikhshift/form"
	"gopkg.in/mgo.v2/bson"
	"path/filepath"
)

const ContentJson string = "application/json"

var Mockname = Mockfilename()

//Path of folder to save session files and 
// load them.
const SessionStorePath string = "."

// Key used to encrypt/decrypt session information. Must
// be a valid AES key.
var Key []byte = []byte("a very very very very secret key")


// Generate mock file name
func Mockfilename() * string {
	name := "file.gen"
	return &name
}



// Save and encrypt the specified interface
// with the specified name. The function will
// add prefix SessionStorePath to the name specifed.
// @test
// @case Mockname, bson.M{} @equal nil 
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
// If testing please make sure you used
// func Save(...
// @test
// @case Mockname @equal bson.M{},nil
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

