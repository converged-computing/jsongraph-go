# JSON Graph Format (JGF) (Go structs)

**under development**

This is a simple library that provides go structures for [JSON Graph schema](https://github.com/jsongraph/json-graph-specification) for use in other projects. The submodule directory with the schema is only
required for testing. We currently provide data structures for each of [version 1](https://github.com/jsongraph/json-graph-specification/blob/master/json-graph-schema_v1.json) and [version 2](https://github.com/jsongraph/json-graph-specification/blob/master/json-graph-schema_v2.json).

## Usage

Build the examples:

```bash
make
```

### Version 2

Run an example - first, cars is a List graph example:

```bash
./examples/v2/bin/cars
```
```console
This example reads in a cars graph
Graph with 4 nodes and 2 edges.
Graph with 3 nodes and 2 edges.
```

Les Miserables (single Graph example)

```bash
./examples/v2/bin/miserables
```
```
This example reads in a les miserables graph
Graph with 77 nodes and 254 edges.
```

This shows a Directed graph (with hyper directed edge type):

```bash
./examples/v2/bin/hyper-directed
```
```console
This example reads in a hyper-directed graph
Graph with 8 nodes and 4 edges.
```

And an undirected type.

```bash
./examples/v2/bin/hyper-undirected
```
```console
This example reads in a hyper-undirected graph
Graph with 6 nodes and 4 edges.
```

And usual suspects!


```bash
./examples/v2/bin/usual-suspects
```
```console
This example reads in a usual suspects graph with metadata
Graph with 2 nodes and 1 edges.
```

### Version 1

We just have one [example from flux-sched](https://github.com/flux-framework/flux-sched/blob/fe872c8dc056934e4073b5fb2932335bb69ca73a/t/data/resource/jgfs/tiny.json):

```bash
./examples/v1/bin/tiny
```
```console
This example reads in tiny v1 graph
Graph with 100 nodes and 198 edges.
```

Note that this library is under development, and we are keeping simple for now! Likely most functionality that you
want we expect you to implement, however if there are common (shared) needs we can add more functions to the structs here.

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614
