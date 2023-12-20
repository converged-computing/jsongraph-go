package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/converged-computing/jsongraph-go/jsongraph/graph"
)

func main() {
	fmt.Println("This example reads in a cars graph")

	// Assumes running from the root
	jsonFileName := flag.String("json", "examples/cars/car_graphs.json", "car_graphs.json file")
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

	data := graph.JsonGraphList{}
	err = json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Printf("error unmarshalling %s:%s\n", jsonFile, err)
		os.Exit(1)
	}
	for _, g := range data.Graphs {
		fmt.Printf("Graph with %d nodes and %d edges.\n", len(g.Nodes), len(g.Edges))
	}
}
