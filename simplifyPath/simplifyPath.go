package simplifypath

import "strings"

func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	stack := []string{}
	var res string
	for _, v := range strs {
		switch v {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
		default:
			if v != "" {
				stack = append(stack, v)
			}
		}
	}

	for _, v := range stack {
		res += "/" + v
	}

	if res == "" {
		return "/"
	}

	return res
}
