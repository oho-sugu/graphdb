package main

import (
	"fmt"

	"strings"

	"github.com/oho-sugu/graphdb"
)

func main() {
	db, _ := graphdb.Open("celegans.db")

	start := db.GetNode(NEURON, no2ID("0"))
	fmt.Println(graphdb.Byte2string(start.ID))

	rec(db, start, 0, 4)
}

func rec(db *graphdb.GraphDB, node *graphdb.Node, curdepth int, maxdepth int) {
	curdepth++
	if curdepth == maxdepth {
		return
	}

	edges := db.GetNodesEdge(node)

	for _, edge := range edges {
		fmt.Println(strings.Repeat(" ", curdepth) + graphdb.Byte2string(edge.To))

		next := db.GetNode(NEURON, edge.To)

		rec(db, next, curdepth, maxdepth)
	}
}
