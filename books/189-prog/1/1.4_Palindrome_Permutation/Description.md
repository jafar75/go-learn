# Palindrom Permutation

Given a string, the task is to find out if the string or any of its permutations are palindrome or not. 

Whitespaces should be ignored.

Trivial solution is that, when there is at most one character in string with odd number of repeats, there is a palindrome permutation.


## Update

There is possible to use a single integer for tracking about at most 1 character with odd repeat. because we don't care about the repeat amount itself and only odd/even is important. So, using an integer with each of its bits is for a-z is sufficient.