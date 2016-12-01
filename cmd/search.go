package main

import (
	"fmt"

	"github.com/oho-sugu/graphdb"
)

func main() {
	db, _ := graphdb.Open("celegans.db")

	start := db.GetNode(NEURON, no2ID("0"))

	edges := db.GetNodesEdge(start)

	for _, edge := range edges {
		fmt.Println(graphdb.Byte2string(edge.To))
	}
}
