package securechain

import (
	"github.com/go-test/deep"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSave(t *testing.T) {

	expgen00 := Save(Mockname, bson.M{})
	if diff := deep.Equal(expgen00, nil); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkSaveMocknamebsonM(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		Save(Mockname, bson.M{})
	}
}

func TestLoad(t *testing.T) {

	expgen00, expgen01 := Load(Mockname)
	if diff := deep.Equal(expgen00, bson.M{}); diff != nil {
		t.Error(diff)
	}

	if diff := deep.Equal(expgen01, nil); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkLoadMockname(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		Load(Mockname)
	}
}

func TestToString(t *testing.T) {

	expgen00 := ToString(bson.M{})
	if diff := deep.Equal(expgen00, "{}"); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkToStringbsonM(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		ToString(bson.M{})
	}
}

func TestToMap(t *testing.T) {

	expgen00, expgen01 := ToMap("{}")
	if diff := deep.Equal(expgen00, bson.M{}); diff != nil {
		t.Error(diff)
	}

	if diff := deep.Equal(expgen01, nil); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkToMap(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		ToMap("{}")
	}
}

func TestLog(t *testing.T) {

	expgen00 := Log([]byte("{}"))
	if diff := deep.Equal(expgen00, bson.M{}); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkLogbyte(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		Log([]byte("{}"))
	}
}

func TestReadBody(t *testing.T) {

	expgen00 := ReadBody(&http.Response{Body: ioutil.NopCloser(strings.NewReader("test"))})
	if diff := deep.Equal(expgen00, []byte("test")); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkReadBodyhttpResponseBodyioutilNopCloserstringsNewReadertest(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		ReadBody(&http.Response{Body: ioutil.NopCloser(strings.NewReader("test"))})
	}
}

func TestEndpoint(t *testing.T) {

	expgen00 := Endpoint("Login")
	if diff := deep.Equal(expgen00, "https://sc.gophersauce.com/momentum/funcs?name=Login"); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkEndpointLogin(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		Endpoint("Login")
	}
}

func TestWrapReq(t *testing.T) {

	expgen00 := WrapReq(bson.M{})
	if diff := deep.Equal(expgen00, bson.M{"req": bson.M{}}); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkWrapReqbsonM(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		WrapReq(bson.M{})
	}
}
