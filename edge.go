package graphdb

// Edge represent graph's edge
type Edge struct {
	From     []byte // must be 4byte
	Dir      int8
	Edgetype int16
	To       []byte // must be 4byte
	Value    []byte
}

// GRAPH DIRECTION CONST
const (
	FORWARD  int8 = 1
	BACKWARD int8 = 2
)

// edgeKeyBytes return bytes array represent Edge's key
func edgeKeyBytes(edge *Edge) []byte {
	var prefixbyte [1]byte
	prefixbyte[0] = EDGE

	ret := append(prefixbyte[:], edge.From...)
	ret = append(ret, byte(edge.Dir),
		byte((edge.Edgetype>>8)&0xff),
		byte(edge.Edgetype&0xff))

	ret = append(ret, edge.To...)

	return ret
}

// edgeKeyBytes return bytes array represent Edge's key with From HASH
func edgeKeyBytesFrom(from []byte) []byte {
	var prefixbyte [1]byte
	prefixbyte[0] = EDGE

	ret := append(prefixbyte[:], from...)

	return ret
}

// edgeKeyBytes return bytes array represent Edge's key with From HASH and Direction
func edgeKeyBytesFromDir(from []byte, dir int8) []byte {
	var prefixbyte [1]byte
	prefixbyte[0] = EDGE

	ret := append(prefixbyte[:], from...)
	ret = append(ret[:], byte(dir))

	return ret
}

// edgeKeyBytes return bytes array represent Edge's key with From HASH and Direction and Edgge Type
func edgeKeyBytesFromDirType(from []byte, dir int8, etype int16) []byte {
	var prefixbyte [1]byte
	prefixbyte[0] = EDGE

	ret := append(prefixbyte[:], from...)
	ret = append(ret[:], byte(dir),
		byte((etype>>8)&0xff),
		byte(etype&0xff))

	return ret
}

func bytesEdge(k []byte, v []byte) *Edge {
	edge := &Edge{}

	key := make([]byte, len(k))
	copy(key, k)

	val := make([]byte, len(v))
	copy(val, v)

	edge.Dir = int8(key[5])
	edge.Edgetype = int16(key[6])<<8 + int16(key[7])
	edge.From = key[1:5]
	edge.To = key[8:12]
	edge.Value = val

	return edge
}
