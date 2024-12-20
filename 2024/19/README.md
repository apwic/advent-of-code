## Day 19

This problem can be simplified to, given some pattern, determine if a string
can be made from that pattern. Based on that my first attemp was to use brute
force. Check for each substring of the pattern, but check from the longsest
first. But, this doesn't work for the input case, the answer was too low.

Based on the discussion in the forum, there are similar leetcode problem to
this which is 139 - Word Break. It appears to be DP problem, for the part 1
the DP that I was using was this:

```
for every pattern, length = n
dp[i] = (s[i-n+1:i] == pattern AND dp[i-n] == True)
```

This formula made sure to check for the length of the pattern to the substring
and check for the word break, which is the `dp[i-n]` part. And this works!

For the part 2, I need to check for the all combination that can be made for
each string that was given. This requires some tweaking to the DP formula. As
my approach for part 1 was bottom up, but for the part 2, I need to use bottom-
up (at least in my understanding). The logic was quite similar, using memoiza-
tion for this, to check for substring from the end of the string.
