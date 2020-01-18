package decoder

var reg map[string]Decoder

func RegisterAlias(name string, dec Decoder) {
    if reg == nil {
        reg = make(map[string]Decoder)
    }
    reg[name] = dec
}

func Register(dec Decoder) {
    RegisterAlias(dec.Name(), dec)
}

func Get(name string) (dec Decoder, ok bool) {
    dec, ok = reg[name]
    return
}
