package blockchain

import (
	"bytes"
	"encoding/binary"
)

func toHex(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}
