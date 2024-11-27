## Question
```
Q2. Write a program to find all n-digit numbers with an equal sum where n varies
    from 1 to 9 and sum <= 81 (Maximum possible sun in a 9-digit number).
    Eg: 5–digit numbers with sum 42 are
    69999 78999 79899 79989 79998 87999 88899 88989 88998 89799 89889 89898 89979 89988 89997
    96999 97899 97989 97998 98799 98889 98898 98979 98988 98997 99699 99789 99798 99879 99888
    99897 99969 99978 99987 99996
```

## Answer
### Explanation of the Program
```
This program generates all n-digit numbers whose digits add up to a given target sum using a 
recursive backtracking approach. It ensures that numbers meet the constraints specified 
(e.g., n-digit, valid starting digit, and target sum).
```

### Key Points of the Code

1. Input Parameters
    * n: Number of digits in the numbers to be generated.
    * target: Sum of the digits for each n-digit number.

2. Recursive Function: findNumberWithSum
The function works by building numbers digit by digit:
Condition 1: if n > 0 && target >= 0
    * Purpose: Ensure the number has not reached n-digits yet and the remaining target sum is valid.
    * Logic:
        * Starts with a valid range of digits:
            * The first digit must not be 0. This is checked using:
                ``` go
                if len(result) == 0 {
                    startDigit = '1'
                }
                ```
              For subsequent digits, the valid range is 0 to 9.
        * Loops through all valid digits ('0' to '9'), appending each to the current result.
        * Recursively calls the function:
            findNumberWithSum(result+string(d), n-1, target-int(d-'0'))
            * Appends the current digit (d) to result. 
            * Decreases n by 1 (fewer digits left to generate).
            * Reduces the target by the numeric value of the current digit (int(d-'0')).

Condition 2: else if n == 0 && target == 0
    * Purpose: Base case to terminate recursion.
    * Logic:
        * When n=0 (all digits have been added) and the remaining sum (target) is also 0, the 
          current number (result) satisfies all conditions.
        * Prints the number:
            ``` go
            fmt.Print(result + " ")
            ```
