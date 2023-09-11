package hashmap

// 13. Roman to Integer
// https://leetcode.com/problems/roman-to-integer
//
// O(n) time | O(1) space - where n is the length of the roman number
func romanToInt(s string) int {
	values := map[int]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var result int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && values[int(s[i])] < values[int(s[i+1])] {
			result -= values[int(s[i])]
		} else {
			result += values[int(s[i])]
		}
	}
	return result
}
