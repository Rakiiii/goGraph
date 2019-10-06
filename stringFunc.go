package graphlib

func countNumbers(str string) int {

	counter := 1
	for _, s := range str {
		if s == ' ' {
			counter++
		}
	}

	return counter
}
