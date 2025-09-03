package scoreastring

/*3110. Score of a String
You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.

Return the score of s.

Example 1:

Input: s = "hello"

Output: 13

Explanation:

The ASCII values of the characters in s are: 'h' = 104, 'e' = 101, 'l' = 108, 'o' = 111. So, the score of s would be |104 - 101| + |101 - 108| + |108 - 108| + |108 - 111| = 3 + 7 + 0 + 3 = 13.

Example 2:

Input: s = "zaz"

Output: 50

Explanation:
The ASCII values of the characters in s are: 'z' = 122, 'a' = 97. So, the score of s would be |122 - 97| + |97 - 122| = 25 + 25 = 50.

Constraints:

2 <= s.length <= 100
s consists only of lowercase English letters.
*/

func scoreOfString(s string) int {
    result := 0

    for i := 0; i < len(s) - 1; i++ {
        diff := int(s[i]) - int(s[i+1])
        if diff < 0 {
            diff = diff * -1
        }

        result += int(diff)
    }

    return result
}

// func main() {
// 	fmt.Println(scoreOfString("hello")) // 13
// 	fmt.Println(scoreOfString("zaz"))   // 50
// }