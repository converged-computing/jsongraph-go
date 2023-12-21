package graph

import (
	"github.com/converged-computing/jsongraph-go/jsongraph/metadata"
)

/*

A JSONGraphSchema can take one of two types:

 - Graphs: a list of graphs
 - Graph: a single graph (map)

Each type can have the same shared edges, nodes and attributes. Note that I considered
using an interface for an abstract type of graph, but thought it would be easier to
represent each type (edge, hyperedge directed/undirected) separately. Then the
user can explicitly ask for what they need without worrying about the interface.

*/

// The graph base has attributes shared across graphs
type GraphBase struct {
	Id       string          `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id,omitempty"`
	Label    string          `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Directed bool            `json:"directed,omitempty" yaml:"directed,omitempty" mapstructure:"directed,omitempty"`
	Type     string          `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Nodes    map[string]Node `json:"nodes" yaml:"nodes" mapstructure:"nodes"`
}

// A JSONGraph can either be parsed into a list of graps or single graph
// We have types for standard, and then explicitly directed and undirected
type JsonGraph struct {
	Graph Graph `json:"graph"`
}

type JsonGraphList struct {
	Graphs []Graph `json:"graphs"`
}

// DirectedJsonGraph is explicitly a directed graph
type DirectedJsonGraph struct {
	Graph DirectedGraph `json:"graph"`
}

type DirectedJsonGraphList struct {
	Graphs []DirectedGraph `json:"graphs"`
}

// UndirectedJsonGraph is explicitly an undirected graph
type UndirectedJsonGraph struct {
	Graph UndirectedGraph `json:"graph"`
}

type UndirectedJsonGraphList struct {
	Graphs []UndirectedGraph `json:"graphs"`
}

// New functions handle creation of new graphs (and directed, etc.)
func NewGraph() *JsonGraph {
	edges := []Edge{}
	nodes := map[string]Node{}
	base := GraphBase{Directed: true, Nodes: nodes}
	return &JsonGraph{Graph: Graph{GraphBase: base, Edges: edges}}
}

func NewDirectedGraph() *DirectedJsonGraph {
	edges := []DirectedEdge{}
	nodes := map[string]Node{}
	base := GraphBase{Directed: true, Nodes: nodes}
	return &DirectedJsonGraph{Graph: DirectedGraph{GraphBase: base, Edges: edges}}
}

func NewUndirectedGraph() *UndirectedJsonGraph {
	edges := []UndirectedEdge{}
	nodes := map[string]Node{}
	base := GraphBase{Directed: false, Nodes: nodes}
	return &UndirectedJsonGraph{Graph: UndirectedGraph{GraphBase: base, Edges: edges}}
}

// A standard Json graph
// - has a list of edges
type Graph struct {
	GraphBase
	Edges []Edge `json:"edges" yaml:"edges" mapstructure:"edges"`
}

// A directed graph
// - has a list of directed edges
type DirectedGraph struct {
	GraphBase
	Edges []DirectedEdge `json:"hyperedges" yaml:"hyperedges" mapstructure:"hyperedges"`
}

// An undirected graph
// - has a list of undirected edges
type UndirectedGraph struct {
	GraphBase
	Edges []UndirectedEdge `json:"hyperedges" yaml:"hyperedges" mapstructure:"hyperedges"`
}

type Node struct {
	Label    *string           `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Metadata metadata.Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`
}

type Edge struct {
	Id       string            `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id,omitempty"`
	Source   string            `json:"source" yaml:"source" mapstructure:"source"`
	Target   string            `json:"target" yaml:"target" mapstructure:"target"`
	Relation string            `json:"relation,omitempty" yaml:"relation,omitempty" mapstructure:"relation,omitempty"`
	Label    string            `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Directed bool              `json:"directed,omitempty" yaml:"directed,omitempty" mapstructure:"directed,omitempty"`
	Metadata metadata.Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`
}

type DirectedEdge struct {
	Id       string            `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id,omitempty"`
	Source   []string          `json:"source" yaml:"source" mapstructure:"source"`
	Target   []string          `json:"target" yaml:"target" mapstructure:"target"`
	Relation string            `json:"relation,omitempty" yaml:"relation,omitempty" mapstructure:"relation,omitempty"`
	Label    string            `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Metadata metadata.Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`
}

type UndirectedEdge struct {
	Id       string            `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id,omitempty"`
	Nodes    []string          `json:"nodes" yaml:"nodes" mapstructure:"nodes"`
	Relation string            `json:"relation,omitempty" yaml:"relation,omitempty" mapstructure:"relation,omitempty"`
	Label    string            `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Metadata metadata.Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`
}
