## Day 22: Monkey Market

### Part 1

This puzzle is about _psuedorandom_ sequence that have pattern for the next
number. The pattern for this next number is determined by some sequence that's
involving the multiplying, dividing, XOR, and modulo. But, the key takeaway
from this sequence is that all the sequence will be modulo by 2^24 in the end.

Part 1 required to generate the 2000 sequence of the input and there's a lot of
input, so just using simple cache is enough. The key of the cache is the number
and the value is the number generated from the sequence.

### Part 2

For this part 2, we only care about the last digit of the secrets. From the
last digit, we need to compare with the digit before and keep in mind the diff
of the previous last digit to the current last digit.

For example, the sequence 123 will produce this last digit:

```
3       \\ from 123
0 (-3)
6 (6)
5 (-1)
4 (-1)
4 (0)
6 (2)
4 (-2)
4 (0)
2 (-2)
```

Based on the difference between the last digit, find the 4 diff sequence that
will produce the highest number from all the secrets. Based on the example
the sequence `-1,-1,0,2` will produce the digit 6 that's the highest among the
all the digit in that current secret.

My approach is to track all the sequence in a map and use the sequence as the
key. The sequence value will contain another map to an original secret number
so that it will count only the first occurence. With this, we can easily track
the sequence that will produce the highest digit among all the secrets number.
But, this approach take around ~2 seconds in my machine.
