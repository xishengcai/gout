package export

import (
	"github.com/xishengcai/gout/dataflow"
)

var (
	defaultCurl = Curl{}
)

func init() {
	dataflow.Register("curl", &defaultCurl)
}
