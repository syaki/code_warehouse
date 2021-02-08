package main

func repeat(character string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return
}
