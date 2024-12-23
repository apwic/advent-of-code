## Day 21: Keypad Conundrum

This puzzle's input is a numeric and a single letter `A`. We need to find the
shortes possible path to create this numeric number. Given the keypad is like
this:

For numeric:

```
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+
```

For directional:

```
    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
```

The twist is that given an input, after we find the path that is needed to
the input. There are several layer to that. We need to find the shortest path
again to the output and several times depending on the layer that we want.

At first, I thought that this is just a graph problem, since we can simulate
the neighbor of the move performed by the keystroke. For example from `0`, we
can go to `2` by direction `^` or `A` by direction `>`. I tried using BFS, but
it didn't produce the shortes possible path. And then, tried using A\* with the
cost is the length of the path (will prioritize the shorter path). This could
work if there's no several layer on top of that.

Then, I tried to hardcode the optimized path from each of the keypad to all the
possible keypad. Since, it's only a small amount of pad it doesn't take much
effort. But, I noticed that I need to prioritize some pad. The key `<` is far
from `A`, it will not give the shortest path for the next layer. So there are
prioritization for the key distance to the key `A`. AND, we need to prioritze
the length of the recurring key. For example from key `2` to `7` there are
several shortest path, one of them which is `^^<` and `^<^`. By choosing the
`^^<` it will yield `AAv<A` for the next layer, but for the `^<^` it will yield
`Av<A>^A`.

For the part 1, by generating a path from the map of the optimal keystroke is
enough. But, for the part 2, it will take too long since generating the path
will be much bigger for each of the layer added. So, after researching the
forum, I handle this by using recursively handling each of the key rather than
processing the input code as a whole. By handling chunk by chunk, we can add
caching to this too. And that works nicely!
