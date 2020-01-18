package decoder

import (
    "io"
    "io/ioutil"
)

func init() {
    RegisterAlias("raw", Raw{})
    RegisterAlias("binary", Raw{})
}

type Raw struct {}

func (d Raw) Name() string {
    return "raw binary"
}

func (d Raw) Decode(input io.Reader) (data []byte, err error) {
    return ioutil.ReadAll(input)
}
