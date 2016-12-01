package main

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/oho-sugu/graphdb"
)

func main() {
	db, _ := graphdb.Open("celegans.db")
	inputFromTSV(db, "c.elegans_neural.male_node.tsv", NEURON, "c.elegans_neural.male_edge.tsv", CONNECTION)

	db.PrintGraph2DOT()
}

// InputFromTSV read two tsv files (node and edge) and insert node and edge into DB
func inputFromTSV(db *graphdb.GraphDB, nodefilename string, nodetype int16, edgefilename string, edgetype int16) {
	nodefile, err := os.Open(nodefilename)
	if err != nil {
		panic(err)
	}
	defer nodefile.Close()

	nodereader := csv.NewReader(nodefile)
	nodereader.Comma = '\t'
	for {
		record, err := nodereader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		node := &graphdb.Node{
			Nodetype: nodetype,
			ID:       no2ID(record[0]),
			Value:    []byte(record[1]),
		}
		db.AddNode(node)
	}

	// ==========================================================================

	edgefile, err := os.Open(edgefilename)
	if err != nil {
		panic(err)
	}
	defer edgefile.Close()

	edgereader := csv.NewReader(edgefile)
	edgereader.Comma = '\t'
	for {
		record, err := edgereader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		db.AddEdge(no2ID(record[0]), no2ID(record[1]), edgetype, []byte(record[2]+","+record[3]))
	}
}
