package interpreter

var reg map[string]Interpreter

func RegisterAlias(name string, interp Interpreter) {
    if reg == nil {
        reg = make(map[string]Interpreter)
    }
    reg[name] = interp
}

func Register(interp Interpreter) {
    RegisterAlias(interp.Name(), interp)
}

func Get(name string) (interp Interpreter, ok bool) {
    interp, ok = reg[name]
    return
}
