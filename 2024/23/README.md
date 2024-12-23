## Day 23: LAN Party

### Part 1

Given input of pair that's separated by `-`, for exampel `kh-tc` and `qp-kh`.
Each of the piar represents undirected graph that is connecting two nodes.
We need to find a cyclic component with the size of 3. On top of that, we need
to find in the cyclic component the node that starts with `t`. Quite simple,
just iterating through the neighbor until get back to the starting node.
