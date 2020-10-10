# leetcode
Solved problem from [leetcode](https://www.leetcode.com)

## #1480 Running Sum of 1d Array
### description
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.
Example 1:
```
Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
```

### solution
My solution Big(O) = O(n)

Explain: We only need n-1 loop and start with index = 1 because as you can see we can guarantee that 0 index is the same from input and for the next index we just add value of nums\[i] with nums\[i - 1] you don't need nested loop for sum it from 0...n of loop because previous index is add it before as shown as table below here

| Index   | Value of index  | Previous value | Sum of index  |
| :-----: | :-------------: | -------------: | ------------  |
| 1       | 2               | 1              | 2 + 1 = 3     |
| 2       | 3               | 3              | 3 + 3 = 6     |
| 3       | 4               | 6              | 4 + 6 = 10    |


## #1431. Kids With the Greatest Number of Candies
### Description
Given the array candies and the integer extraCandies, where candies[i] represents the number of candies that the ith kid has.

For each kid check if there is a way to distribute extraCandies among the kids such that he or she can have the greatest number of candies among them. Notice that multiple kids can have the greatest number of candies.

Example 1:
```
Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true] 
Explanation: 
Kid 1 has 2 candies and if he or she receives all extra candies (3) will have 5 candies --- the greatest number of candies among the kids. 
Kid 2 has 3 candies and if he or she receives at least 2 extra candies will have the greatest number of candies among the kids. 
Kid 3 has 5 candies and this is already the greatest number of candies among the kids. 
Kid 4 has 1 candy and even if he or she receives all extra candies will only have 4 candies. 
Kid 5 has 3 candies and if he or she receives at least 2 extra candies will have the greatest number of candies among the kids. 
```

### Solution
My solution Big(O): O(2n)

Explain: We need to find max_number in candies array input and now we just need n loop through array for check candy of each kid has greater than or equal to max_number or not and set it to array of output.

## #1512. Number of Good Pairs
Given an array of integers nums.

A pair (i, j) is called good if nums\[i] == nums\[j] and i < j.

Return the number of good pairs.

Example 1:

```
Input: nums = [1, 2, 3, 1, 1, 3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
```

### Solution
My solution Big(O): O(n)

Explain: Make map\[int]int to contains map\[key] of int value that we found. Then we need to iterate through input and set int found = 1 by default check now we check it from map if we contain key we add pairs, found by get value from map\[key]. After we set found to map\[key] and continue it to n

## #771. Jewels and Stones
You're given strings J representing the types of stones that are jewels, and S representing the stones you have.  Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a" is considered a different type of stone from "A".

Example 1:
```
Input: J = "aA", S = "aAAbbbb"
Output: 3
Example 2:
```
Example 2:
```
Input: J = "z", S = "ZZ"
Output: 0
```

Note:
  - ```S``` and ```J``` will consist of letters and have length at most 50.
  - The characters in J are distinct.
  
### Solution
Explain: defined count for sum of result and split ```S``` to map\["string"]int, then iterate through ```J``` and check it is key from ``` map[ S[n] ]``` if it true and plus count by 1 after iterate just return sum of count.