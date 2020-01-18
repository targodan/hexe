package interpreter

import "github.com/targodan/hexe/endianess"

type Interpretation struct {
    Description string
    Value string
    Endianess endianess.T
}

type Interpreter interface {
    Name() string
    Interpret(data []byte) (intData []*Interpretation, err error)
}
