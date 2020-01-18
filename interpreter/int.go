package interpreter

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"github.com/targodan/hexe/endianess"
)

func init() {
	Register(Int64Little{})
	Register(Int64Big{})
	Register(Joined("sint64", Int64Little{}, Int64Big{}))
	Register(UInt64Little{})
	Register(UInt64Big{})
	Register(Joined("uint64", UInt64Little{}, UInt64Big{}))
	Register(Joined("int64_little", Int64Little{}, UInt64Little{}))
	Register(Joined("int64_big", Int64Big{}, UInt64Big{}))
	Register(Joined("int64", Int64Little{}, UInt64Little{}, Int64Big{}, UInt64Big{}))

	Register(Int32Little{})
	Register(Int32Big{})
	Register(Joined("sint32", Int32Little{}, Int32Big{}))
	Register(UInt32Little{})
	Register(UInt32Big{})
	Register(Joined("uint32", UInt32Little{}, UInt32Big{}))
	Register(Joined("int32_little", Int32Little{}, UInt32Little{}))
	Register(Joined("int32_big", Int32Big{}, UInt32Big{}))
	Register(Joined("int32", Int32Little{}, UInt32Little{}, Int32Big{}, UInt32Big{}))

    Register(Joined("int", Int64Little{}, UInt64Little{}, Int64Big{}, UInt64Big{}, Int32Little{}, UInt32Little{}, Int32Big{}, UInt32Big{}))
}

type Int64Little struct{}

func (in Int64Little) Name() string {
	return "sint64_little"
}

func (in Int64Little) Interpret(data []byte) (intData []*Interpretation, err error) {
	var num int64
	err = binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &num)

	return []*Interpretation{
		&Interpretation{
			Description: "signed int64",
			Value:       strconv.FormatInt(num, 10),
			Endianess:   endianess.Little,
		},
	}, err
}

type Int64Big struct{}

func (in Int64Big) Name() string {
	return "sint64_big"
}

func (in Int64Big) Interpret(data []byte) (intData []*Interpretation, err error) {
	var num int64
	err = binary.Read(bytes.NewBuffer(data), binary.BigEndian, &num)

	return []*Interpretation{
		&Interpretation{
			Description: "signed int64",
			Value:       strconv.FormatInt(num, 10),
			Endianess:   endianess.Big,
		},
	}, err
}

type UInt64Little struct{}

func (in UInt64Little) Name() string {
	return "uint64_little"
}

func (in UInt64Little) Interpret(data []byte) (intData []*Interpretation, err error) {
	var num uint64
	err = binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &num)

	return []*Interpretation{
		&Interpretation{
			Description: "unsigned int64",
			Value:       strconv.FormatUint(num, 10),
			Endianess:   endianess.Little,
		},
	}, err
}

type UInt64Big struct{}

func (in UInt64Big) Name() string {
	return "uint64_big"
}

func (in UInt64Big) Interpret(data []byte) (intData []*Interpretation, err error) {
	var num uint64
	err = binary.Read(bytes.NewBuffer(data), binary.BigEndian, &num)

	return []*Interpretation{
		&Interpretation{
			Description: "unsigned int64",
			Value:       strconv.FormatUint(num, 10),
			Endianess:   endianess.Big,
		},
	}, err
}

type Int32Little struct{}

func (in Int32Little) Name() string {
	return "sint32_little"
}

func (in Int32Little) Interpret(data []byte) (intData []*Interpretation, err error) {
	var num int32
	err = binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &num)

	return []*Interpretation{
		&Interpretation{
			Description: "signed int32",
			Value:       strconv.FormatInt(int64(num), 10),
			Endianess:   endianess.Little,
		},
	}, err
}

type UInt32Little struct{}

func (in UInt32Little) Name() string {
	return "uint32_little"
}

func (in UInt32Little) Interpret(data []byte) (intData []*Interpretation, err error) {
	var num uint32
	err = binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &num)

	return []*Interpretation{
		&Interpretation{
			Description: "unsigned int32",
			Value:       strconv.FormatUint(uint64(num), 10),
			Endianess:   endianess.Little,
		},
	}, err
}

type Int32Big struct{}

func (in Int32Big) Name() string {
	return "sint32_big"
}

func (in Int32Big) Interpret(data []byte) (intData []*Interpretation, err error) {
	var num int32
	err = binary.Read(bytes.NewBuffer(data), binary.BigEndian, &num)

	return []*Interpretation{
		&Interpretation{
			Description: "signed int32",
			Value:       strconv.FormatInt(int64(num), 10),
			Endianess:   endianess.Big,
		},
	}, err
}

type UInt32Big struct{}

func (in UInt32Big) Name() string {
	return "uint32_big"
}

func (in UInt32Big) Interpret(data []byte) (intData []*Interpretation, err error) {
	var num uint32
	err = binary.Read(bytes.NewBuffer(data), binary.BigEndian, &num)

	return []*Interpretation{
		&Interpretation{
			Description: "unsigned int32",
			Value:       strconv.FormatUint(uint64(num), 10),
			Endianess:   endianess.Big,
		},
	}, err
}
