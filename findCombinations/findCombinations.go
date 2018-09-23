package findCombinations

import "strings"

func FindCombinations(s string) chan string {
	ch := make(chan string)
	go func() {
		queue := []string{}
		queue = append(queue, s)
		first := ""
		for len(queue) > 0 {
			first, queue = queue[0], queue[1:]

			if idx := strings.Index(first, "X"); idx != -1 {
				tmp := []rune(first)
				tmp[idx] = '0'
				queue = append(queue, string(tmp))

				tmp[idx] = '1'
				queue = append(queue, string(tmp))
			} else {
				ch <- first
			}
		}
		close(ch)
	}()
	return ch
}
