# TDD

## Discipline
* Write a test.
* Make the compiler pass.
* Run the test, see that it fails and check the error message is meaningful.
* Write enough code to make the test pass.
* Refactor.

## TDD process and why the steps are important
* _Write a failing test and see_ it fails, so we know we have
written a _relevant test_ for our requirements and seen it produces an _easy to 
to understand description of failure_.
*  Writing the smallest amount of code to make it pass so know we have working
software.
* Then refactor, backed with the safety of our tests to ensure 
we have well-crafted code that is easy to refactor with

## Some Steps on How to Write Tests ( Simpler version of above in my words)
* Write Test and it will not compile because the function that 
is used in test does not exist anywhere in code.
* After the compilation fails, write simplest defination of that function 
that will let the program compile, but the test will fail becuase there is no logic
written, it could just return zero.
* After the test fails, write the logic such that that test case passes.
* Iterate by adding more test cases and keep changing the logic. 

## Some Problems that you may encounter
* Sometimes due to change in the requirement, the code written for previous requirement may break.
* For example, if you have written a function that takes in array of length 5
and return sum of all the elements in the array. The usage of this function will require
to pass an array of length 5 only.
* This would result in these two cases that you can do to resolve:-
    * Break the existing API by changing the argument to Sum to be a slice rather
 than an array. When we do this we will know we have potentially ruined
 someone's day because our other test will not compile!

    * Create a new function