package interpreter

import "github.com/targodan/go-errors"

func Joined(name string, subs ...Interpreter) Interpreter {
	return &joined{
		name: name,
		Subs: subs,
	}
}

type joined struct {
	name string
	Subs []Interpreter
}

func (in *joined) Name() string {
	return in.name
}

func (in *joined) Interpret(data []byte) (intData []*Interpretation, err error) {
	interpretations := make([]*Interpretation, 0)

	for _, interp := range in.Subs {
		subInterpretations, subErr := interp.Interpret(data)

		err = errors.NewMultiError(err, subErr)
        if subErr == nil {
            interpretations = append(interpretations, subInterpretations...)
        }
	}
	return interpretations, err
}
