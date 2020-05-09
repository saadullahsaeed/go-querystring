package querystring_test

import (
	"net/url"
	"testing"

	"reflect"

	"github.com/saadullahsaeed/go-querystring"
)

type testStruct struct {
	A string  `url:"a"`
	B uint64  `url:"b"`
	C uint64  `url:"-"`
	D float32 `url:"d"`
	E float64 `url:"e"`
}

func TestDecode(t *testing.T) {
	values := url.Values{}
	values.Add("a", "test")
	values.Add("b", "1")
	values.Add("c", "1")
	values.Add("d", "0.2")
	values.Add("e", "1.7")

	ts := &testStruct{}
	querystring.Decode(values, ts)
	expected := &testStruct{A: "test", B: 1, D: 0.2, E: 1.7}
	if !reflect.DeepEqual(expected, ts) {
		t.FailNow()
	}
}
