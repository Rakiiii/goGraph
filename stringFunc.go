package graph

func countNumbers(str string) int {

	counter := 0
	for _, s := range str {
		if s == ' ' {
			counter++
		}
	}

	return counter
}
