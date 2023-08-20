package common

import "fmt"

func Translate(actions int) string {
	var result string
	for idx := range order {
		if actions&order[idx] == order[idx] {
			result += fmt.Sprintf("%v->", aliases[order[idx]])
		}
	}
	return result
}
