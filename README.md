# Fibonacci Number Sequance #

The Fibonacci numbers are the numbers in the following integer sequence.
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...

In mathematical terms, the sequence Fn of Fibonacci numbers is defined by the recurrence relation

<code>Fn = Fn-1 + Fn-2</code>

with seed values

<code>F0 = 0 and F1 = 1</code>

Given a number n, print n-th Fibonacci Number.

### Examples ###

<code>
Input  : n = 2 -> Output : 1</code>
<br>
<code>Input  : n = 9 -> Output : 34</code>

### How to get set up ###

1. Edit the constant values defined at the top of the main function
   1. <code>workers</code> Is the number of workers that will be setup. This should not exceed the number of cores on the CPU being used otherwise performance will be hindered
   2. <code>limit</code> Is the maximum <code>n</code> number
2. Run the program and see the numbers be printed out
