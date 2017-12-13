package securechain

import (
	"github.com/go-test/deep"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

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

func TestSave(t *testing.T) {

	expgen00 := Save(Mockfilename(), bson.M{})
	if diff := deep.Equal(expgen00, nil); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkSaveMockfilenamebsonM(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		Save(Mockfilename(), bson.M{})
	}
}

func TestLoad(t *testing.T) {

	expgen00, expgen01 := Load(Mockfilename())
	if diff := deep.Equal(expgen00, bson.M{}); diff != nil {
		t.Error(diff)
	}

	if diff := deep.Equal(expgen01, nil); diff != nil {
		t.Error(diff)
	}

}

func BenchmarkLoadMockfilename(b *testing.B) {
	// credits to @davecheney
	for n := 0; n < b.N; n++ {
		Load(Mockfilename())
	}
}
