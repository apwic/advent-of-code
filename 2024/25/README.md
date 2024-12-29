## Day 25: Code Chronicle

### Part 1

Given a grid with size of `6 x 5`, that is either a key or a lock. Key consist
of the top row filled with `#` and the bottom row filled with `.`. Lock is the
opposite of the key. The task is to find all the possible fit between the lock
and the key.

At first, I thought that both the lock and key needed to be exact fit, so using
map to find the combination will be faster. But, it turns out that the fit can
be not exact as long as the lock and the key not overlap. So, map can't be used
here. We need to iterate all the combination between the key and lock to find
all the fit.
