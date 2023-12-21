package array

// 733. Flood Fill
// https://leetcode.com/problems/flood-fill/
//
// O(m * n) time | O(m * n) space
func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	clr := image[sr][sc]
	image[sr][sc] = color
	stack := [][2]int{{sr, sc}}

	for len(stack) > 0 {
		pixel := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		x, y := pixel[0], pixel[1]

		neighbors := [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
		for _, n := range neighbors {
			x1, y1 := n[0], n[1]
			if x1 >= 0 && x1 < len(image) && y1 >= 0 && y1 < len(image[0]) {
				if image[x1][y1] != color && image[x1][y1] == clr {
					stack = append(stack, [2]int{x1, y1})
					image[x1][y1] = color
				}
			}
		}
	}

	return image
}
