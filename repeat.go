package main

func Repeat(character string, repeatCount int) string {
	// The stdlib alternative way to do this is return strings.Repeat(character, repeatCount)

	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
