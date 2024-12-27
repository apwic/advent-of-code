## Day 23: LAN Party

### Part 1

Given input of pair that's separated by `-`, for example `kh-tc` and `qp-kh`.
Each of the pair represents undirected graph that is connecting two nodes.
We need to find a cyclic component with the size of 3. On top of that, we need
to find in the cyclic component the node that starts with `t`. Quite simple,
just iterating through the neighbor until get back to the starting node.

### Part 2

Part 2 is a lot different than the Part 1. In this part 2, we are asked to
find the largest interconnected node. I looked around and apparently this
largest interconnected is called _maximum clique_. The algorithm to find this
maximum clique is called Bron-Kerbosch algorithm. Using that algorithm it is
quite simple to get the answer.
