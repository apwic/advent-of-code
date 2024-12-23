## Day 20: Race Condition

### Part 1

Given a grid which consist of:

- `.` -> empty
- `#` -> wall
- `S` -> start
- `E` -> end

Different than the usual path finding algorithm, this puzzle only have 1 path.
The twist is that you can cheat by moving through the obstacle, the max amount
of cheat that can be used is 2. You need to find the different time that is
gotten by cheating.

Example:

```
. . . # E
. # . # .
. # . # .
S # . . .
```

You can cheat by doing this:

```
. . . # E
. # . # .
. # . # .
S 1 2 . .
```

Doing that cheat will save the cost in the path. My approach is that you need
to check for the all combination of path that you can cheat. This means to ite-
rate through the path that is available and attempt a cheat. When the cheat is
possible then count the different by using the previous cost and the current
cost. Previous cost is cost that is made by following the regular path, and the
current cost is that the cheat cost. This cheat cost is made by adding the
first cost before the cheat start and add it by the obstacle that you pass
through.

## Part 2

The part 2 increased the amount of the cheats allowed from 2 to 20. My imple-
mentation before for part 1 was taking too long, as I iterate through the
direction (Up, Down, Left, Right). Then, I was looking for other approach and
I checked the forum for different approach. I used the term offset which search
for the radius (in this case is the amount of the cheats). Then, I need to
adjust the iteration and how to handle the distance.
