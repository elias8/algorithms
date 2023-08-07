package array

// 74. Search a 2D Matrix
// https://leetcode.com/problems/search-a-2d-matrix
//
// O(log(n*m) time | O(1) space
func searchMatrix(matrix [][]int, target int) bool {
	width, top, bottom := len(matrix[0])-1, 0, len(matrix)-1
	for top <= bottom {
		yMid := (top + bottom) / 2
		if target >= matrix[yMid][0] && target <= matrix[yMid][width] {
			left, right := 0, len(matrix[0])-1
			for left <= right {
				xMid := (left + right) / 2
				if matrix[yMid][xMid] == target {
					return true
				} else if target > matrix[yMid][xMid] {
					left = xMid + 1
				} else {
					right = xMid - 1
				}
			}
			break
		} else if target > matrix[yMid][width] {
			top = yMid + 1
		} else {
			bottom = yMid - 1
		}
	}
	return false
}
