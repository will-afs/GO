package main

func Density(str string) (rets map[rune]int) {
	density := make(map[rune]int)
	for i := 0; i < len(str); i++ {
		char := str[i]
		if val, ok := density[char]; ok {
			density[char]++
		} else {
			density[char] = 0
		}
	}
	return density
}

func main() {
}
