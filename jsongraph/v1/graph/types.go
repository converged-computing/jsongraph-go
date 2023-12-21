package graph

/*

A JSONGraphSchema can take one of two types:

 - Graphs: a list of graphs with metadata at the top leve
 - Graph: a single graph (map)

*/

// A JSONGraph can either be parsed into a list of graps or single graph
// We have types for standard, and then explicitly directed and undirected
type JsonGraph struct {
	Graph Graph `json:"graph"`
}

// Metadata is a basic map that can be used by any graph object
type Metadata map[string]interface{}

type JsonGraphList struct {
	Label    string   `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Type     string   `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Metadata Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`
	Graphs   []Graph  `json:"graphs"`
}

// The base graph type
type Graph struct {
	Label    string   `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Directed bool     `json:"directed,omitempty" yaml:"directed,omitempty" mapstructure:"directed,omitempty"`
	Type     string   `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Metadata Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`
	Nodes    []Node   `json:"nodes" yaml:"nodes" mapstructure:"nodes"`
	Edges    []Edge   `json:"edges,omitempty" yaml:"edges,omitempty" mapstructure:"edges,omitempty"`
}

type Node struct {
	Id       string   `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id,omitempty"`
	Label    *string  `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Metadata Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`
}

type Edge struct {
	Id       string   `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id,omitempty"`
	Source   string   `json:"source" yaml:"source" mapstructure:"source"`
	Target   string   `json:"target" yaml:"target" mapstructure:"target"`
	Relation string   `json:"relation,omitempty" yaml:"relation,omitempty" mapstructure:"relation,omitempty"`
	Label    string   `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`
	Directed bool     `json:"directed,omitempty" yaml:"directed,omitempty" mapstructure:"directed,omitempty"`
	Metadata Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty" mapstructure:"metadata,omitempty"`
}

// New functions handle creation of new graphs (and directed, etc.)
func NewGraph() *JsonGraph {
	edges := []Edge{}
	nodes := []Node{}
	return &JsonGraph{Graph: Graph{Nodes: nodes, Edges: edges, Directed: true}}
}
