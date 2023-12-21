package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/converged-computing/jsongraph-go/jsongraph/v2/graph"
)

func main() {
	fmt.Println("This example reads in a hyper-undirected graph")
	jsonFileName := flag.String("json", "examples/v2/hyper-undirected/hyper-undirected.json", "json input file")
	flag.Parse()

	jsonFile := *jsonFileName
	if jsonFile == "" {
		flag.Usage()
		os.Exit(0)
	}
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("error reading %s:%s\n", jsonFile, err)
		os.Exit(1)
	}

	g := graph.UndirectedJsonGraph{}
	err = json.Unmarshal([]byte(file), &g)
	if err != nil {
		fmt.Printf("error unmarshalling %s:%s\n", jsonFile, err)
		os.Exit(1)
	}

	fmt.Printf("Graph with %d nodes and %d edges.\n", len(g.Graph.Nodes), len(g.Graph.Edges))
}
