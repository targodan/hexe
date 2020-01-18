package decoder

import (
    "io"
    "io/ioutil"
    "encoding/base64"
)

var Base64Encoding = base64.StdEncoding

func init() {
    Register(Base64{})
}

type Base64 struct {}

func (d Base64) Name() string {
    return "base64"
}

func (d Base64) Decode(input io.Reader) (data []byte, err error) {
    dec := base64.NewDecoder(Base64Encoding, input)
    return ioutil.ReadAll(dec)
}
