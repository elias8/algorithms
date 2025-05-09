package string

import "strings"

// 609. Find Duplicate File in System
// https://leetcode.com/problems/find-duplicate-file-in-system
//
// O(n) time | O(n) space: where n is the number of characters
func findDuplicate(paths []string) [][]string {
	contentToFile := map[string][]string{}
	for _, path := range paths {
		files := strings.Split(path, " ")
		directory := files[0]

		for i := 1; i < len(files); i++ {
			part := files[i]
			opening := strings.Index(part, "(")
			closing := strings.Index(part, ")")
			file, content := part[:opening], part[opening:closing]
			contentToFile[content] = append(contentToFile[content], directory+"/"+file)
		}
	}

	var duplicates [][]string
	for _, value := range contentToFile {
		if len(value) > 1 {
			duplicates = append(duplicates, value)
		}
	}
	return duplicates
}
