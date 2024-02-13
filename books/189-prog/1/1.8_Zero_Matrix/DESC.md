# Zero Matrix

Write an algorithm such that if an element in an MxN matrix is 0, its entire row and
column are set to 0.

It seems easy, but some notes should be considered:

a. checking a cell is zero, and then set enitre row and column of it to 0, during iteration may cause issue

b. one can solve the problem with a space of O(m+n)   (m rows and n cols), but solving with constant space i.e. using the space of matrix itself, might be a little tricky