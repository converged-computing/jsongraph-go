package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/converged-computing/jsongraph-go/jsongraph/v1/graph"
)

func main() {
	fmt.Println("This example reads in tiny v1 graph")

	// Assumes running from the root
	jsonFileName := flag.String("json", "examples/v1/tiny/tiny.json", "json file")
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

	g := graph.JsonGraph{}
	err = json.Unmarshal([]byte(file), &g)
	if err != nil {
		fmt.Printf("error unmarshalling %s:%s\n", jsonFile, err)
		os.Exit(1)
	}
	fmt.Printf("Graph with %d nodes and %d edges.\n", len(g.Graph.Nodes), len(g.Graph.Edges))
	_, err = json.Marshal(g.Graph)
	if err != nil {
		fmt.Printf("error marshalling %s:%s\n", jsonFile, err)
		os.Exit(1)
	}

}
