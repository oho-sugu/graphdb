package graphdb

// Node represent graph's node
type Node struct {
	Nodetype int16
	ID       []byte // must be 4byte
	Value    []byte
}

// nodeKeyBytes return bytes array represent Node's key
func nodeKeyBytes(node *Node) []byte {
	var retbyte [3]byte
	retbyte[0] = NODE
	retbyte[1] = byte((node.Nodetype >> 8) & 0xff)
	retbyte[2] = byte(node.Nodetype & 0xff)

	ret := append(retbyte[:], node.ID[:]...)

	return ret
}

// nodeKeyBytesType return bytes array represent Node's key only type prefix
func nodeKeyBytesType(nodetype int16) []byte {
	var retbyte [3]byte
	retbyte[0] = NODE
	retbyte[1] = byte((nodetype >> 8) & 0xff)
	retbyte[2] = byte(nodetype & 0xff)
	return retbyte[:]
}

func bytesNode(k []byte, v []byte) *Node {
	node := &Node{}

	key := make([]byte, len(k))
	copy(key, k)

	val := make([]byte, len(v))
	copy(val, v)

	node.Nodetype = int16(key[1])<<8 + int16(key[2])
	node.ID = key[3:7]
	node.Value = val

	return node
}
