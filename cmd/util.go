package main

import (
	"encoding/binary"
	"strconv"
)

func no2ID(no string) []byte {
	noint, _ := strconv.Atoi(no)
	id := make([]byte, 4)
	binary.LittleEndian.PutUint32(id, uint32(noint))
	return id
}
