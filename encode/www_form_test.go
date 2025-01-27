package encode

import (
	"bytes"
	"testing"

	"github.com/xishengcai/gout/core"
	"github.com/xishengcai/gout/setting"
	"github.com/stretchr/testify/assert"
)

type testWWWForm struct {
	w    *WWWFormEncode
	in   interface{}
	need string
}

func Test_WWWForm_Encode(t *testing.T) {
	var out bytes.Buffer

	tests := []testWWWForm{
		{NewWWWFormEncode(setting.Setting{}), core.A{"k1", "v1", "k2", 2, "k3", 3.14}, "k1=v1&k2=2&k3=3.14"},
		{NewWWWFormEncode(setting.Setting{}), core.H{"k1": "v1", "k2": 2, "k3": 3.14}, "k1=v1&k2=2&k3=3.14"},
	}

	for _, v := range tests {
		assert.NoError(t, v.w.Encode(v.in))
		assert.NoError(t, v.w.End(&out))
		assert.Equal(t, out.String(), v.need)
		out.Reset()
	}

}

func Test_WWWForm_Name(t *testing.T) {
	assert.Equal(t, NewWWWFormEncode(setting.Setting{}).Name(), "www-form")
}
