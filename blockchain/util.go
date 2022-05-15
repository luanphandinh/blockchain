package blockchain

import (
	"bytes"
	"encoding/binary"
)

// @TODO: return error here
func toHex(num int64) []byte {
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		// bad practice
		panic(err)
	}

	return buff.Bytes()
}
