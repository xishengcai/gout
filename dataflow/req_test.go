package dataflow

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/xishengcai/gout/encode"
	"github.com/stretchr/testify/assert"
)

func TestReqModifyUrl(t *testing.T) {
	src := []string{"127.0.0.1", ":8080/query", "/query", "http://127.0.0.1", "https://127.0.0.1"}
	want := []string{"http://127.0.0.1", "http://127.0.0.1:8080/query", "http://127.0.0.1/query", "http://127.0.0.1", "https://127.0.0.1"}

	for k, v := range src {
		if want[k] != modifyURL(v) {
			t.Errorf("got %s want %s\n", modifyURL(v), want[k])
		}
	}
}

type urlTest struct {
	set  interface{}
	need interface{}
}

func TestReq_isAndGetString(t *testing.T) {
	test := []urlTest{
		{set: "?a=b&c=d", need: "a=b&c=d"},
		{set: "c=d&e=f", need: "c=d&e=f"},
		{set: []byte("c=d&e=f"), need: "c=d&e=f"},
		{set: time.Time{}, need: ""},
	}

	for _, v := range test {
		rv, _ := isAndGetString(v.set)
		assert.Equal(t, v.need, rv)
	}
}

// 测试request()函数调用出错的情况
func TestReq_request_fail(t *testing.T) {

	tests := []func() *Req{
		func() *Req {
			r := Req{}
			r.bodyEncoder = encode.NewBodyEncode(&map[string]string{})
			return &r
		},
		func() *Req {
			r := Req{}
			s := "hello"
			r.form = []interface{}{s}
			return &r
		},
	}

	for _, test := range tests {
		r := test()
		_, err := r.Request()
		assert.Error(t, err)
	}

}

type testValid struct {
	Val string `valid:"required"`
}

func Test_Valid(t *testing.T) {
	total := int32(1)
	router := setupMethod(&total)

	ts := httptest.NewServer(http.HandlerFunc(router.ServeHTTP))
	defer ts.Close()

	testCases := []string{"bindheader"}
	for _, c := range testCases {
		val := testValid{}
		g := GET(ts.URL + "/someGet")
		var err error
		switch c {
		case "bindheader":
			err = g.BindHeader(&val).Do()
		}

		//fmt.Printf("-->%v\n", err)
		assert.Error(t, err)
	}
}
