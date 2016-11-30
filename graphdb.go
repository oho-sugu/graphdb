package graphdb

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// GraphDB is a GraphDB main struct
type GraphDB struct {
	Leveldb *leveldb.DB
}

// KeyPrefixes ...
const (
	NODE byte = 0x01
	EDGE byte = 0x02
)

// Open open GraphDB in path
func Open(path string) (db *GraphDB, err error) {
	db = &GraphDB{}
	opt := &opt.Options{
		ErrorIfExist: false,
	}
	db.Leveldb, err = leveldb.OpenFile(path, opt)
	if err != nil {
		panic("Error opening LevelDB File : " + path + "\n" + err.Error())
	}

	return
}

// AddNode add node data to db
func (db *GraphDB) AddNode(node *Node) {
	db.Leveldb.Put(nodeKeyBytes(node), node.Value, nil)
}

// GetNode get node with type and id
func (db *GraphDB) GetNode(nodetype int16, id []byte) *Node {
	node := &Node{}
	node.Nodetype = nodetype
	node.ID = id

	bytes, err := db.Leveldb.Get(nodeKeyBytes(node), nil)
	if bytes == nil {
		return nil
	}
	if err != nil {
		panic(err)
	}
	node.Value = bytes
	return node
}

// _addEdge add edge data to db (internal use only)
func (db *GraphDB) _addEdge(edge *Edge) {
	db.Leveldb.Put(edgeKeyBytes(edge), edge.Value, nil)
}

// AddEdge add bidiretional edge to db
func (db *GraphDB) AddEdge(from []byte, to []byte, etype int16, value []byte) {
	edge := &Edge{}
	edge.Dir = FORWARD
	edge.Edgetype = etype
	edge.From = from
	edge.To = to
	edge.Value = value
	db._addEdge(edge)

	edge.From = to
	edge.To = from
	edge.Dir = BACKWARD
	db._addEdge(edge)

}

// GetNodesEdge get all edge associates node
func (db *GraphDB) GetNodesEdge(node *Node) []*Edge {
	var edges []*Edge

	it := db.Leveldb.NewIterator(util.BytesPrefix(edgeKeyBytesFrom(node.ID)), nil)

	for it.Next() {
		edge := bytesEdge(it.Key(), it.Value())
		edges = append(edges, edge)
	}

	return edges
}

// GetNodesByType get all nodes whitch type is "nodetype"
func (db *GraphDB) GetNodesByType(nodetype int16) []*Node {
	var nodes []*Node

	it := db.Leveldb.NewIterator(util.BytesPrefix(nodeKeyBytesType(nodetype)), nil)

	for it.Next() {
		node := bytesNode(it.Key(), it.Value())
		nodes = append(nodes, node)
	}

	return nodes
}

func (db *GraphDB) getWholeNodes() []*Node {
	var nodes []*Node

	it := db.Leveldb.NewIterator(util.BytesPrefix([]byte{NODE}), nil)
	for it.Next() {
		node := bytesNode(it.Key(), it.Value())
		nodes = append(nodes, node)
	}

	return nodes
}

func (db *GraphDB) getWholeEdges() []*Edge {
	var edges []*Edge

	it := db.Leveldb.NewIterator(util.BytesPrefix([]byte{EDGE}), nil)
	for it.Next() {
		edge := bytesEdge(it.Key(), it.Value())
		edges = append(edges, edge)
	}

	return edges
}

// ============ Some Utility Function for Graph Visualize and Graph Creation ======

// PrintGraph2DOT print out whole graph to simple DOT format
func (db *GraphDB) PrintGraph2DOT() {
	fmt.Println("digraph negraph {")

	edges := db.getWholeEdges()

	for _, edge := range edges {
		if edge.Dir == FORWARD {
			fmt.Printf("  %s -> %s;\n", byte2string(edge.From), byte2string(edge.To))
		}
	}

	fmt.Println("}")
}

func byte2string(val []byte) string {
	var str string

	for _, b := range val {
		str = str + fmt.Sprintf("%d", b)
	}

	return "NODE_" + str
}
