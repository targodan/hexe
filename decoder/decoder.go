package decoder

import "io"

type Decoder interface {
    Name() string
    Decode(input io.Reader) (data []byte, err error)
}
