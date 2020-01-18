// Package endianess contains information on how eat eggs.
package endianess

import (
    "strings"
    "errors"
)

// T holds information about endianess
type T int

const (
    // Independent means, the data is independent of the source
    // or current endianess. UTF-8 for example is by definition independant
    // of the architechtures endianess.
    Independent T = iota
    // Little means the data is interpreted as little endian.
    Little
    // Big means, the data is interpreted as big endian.
    Big
)

func (e T) String() string {
    switch e {
    case Independent:
        return "independent of endianess"
    case Little:
        return "little endian"
    case Big:
        return "big endian"
    }
    return "INVALID ENDIANESS"
}

func FromString(str string) (T, error) {
    str = strings.ToLower(str)

    littleKeyword := "little"
    bigKeyword := "big"
    independenKeyword := "independent"

    if len(str) >= len(littleKeyword) && str[0:len(littleKeyword)] == littleKeyword {
        return Little, nil
    }
    if len(str) >= len(bigKeyword) && str[0:len(bigKeyword)] == bigKeyword {
        return Big, nil
    }
    if len(str) >= len(independenKeyword) && str[0:len(independenKeyword)] == independenKeyword {
        return Independent, nil
    }

    return Independent, errors.New("invalid endianess")
}
