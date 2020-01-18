package decoder

import (
    "io"
    "io/ioutil"
    "encoding/hex"
)

func init() {
    Register(Hex{})
}

type Hex struct {}

func (d Hex) Name() string {
    return "hex"
}

func (d Hex) Decode(input io.Reader) (data []byte, err error) {
    dec := hex.NewDecoder(input)
    return ioutil.ReadAll(dec)
}
