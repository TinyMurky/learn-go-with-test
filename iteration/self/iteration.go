// Package iteration a
package iteration

import "strings"

// 講稿
// 1. return ""
// 2. c type for 1 ~ 5
// 3. while 寫法
// 4. range 寫法
// 5. 無限迴圈寫法
// 6. Benchmark
// 7. string builder
// 8. specitic how many times
// 9. example

// Repeat1 c type for
func Repeat1(character string) string {
	var repeated string

	for i := 0; i < 5; i++ { // i 沒有括號
		repeated += character
	}

	return repeated
}

// Repeat2 while 寫法
func Repeat2(character string) string {
	var repeated string

	i := 0
	for i < 5 { // i 沒有括號
		repeated += character
		i++
	}

	return repeated
}

// Repeat3 range 寫法
func Repeat3(character string) string {
	var repeated string

	// for i := range 5 {
	// 	repeated += character
	// 	fmt.Println(i)
	// }

	for range 5 {
		repeated += character
	}

	return repeated
}

// Repeat4 forever
// func Repeat4(character string) string {
// 	var repeated string
//
// 	for {
// 		repeated += character
// 	}
//
// 	return repeated
// }

// Repeat5 用 string builder
func Repeat5(character string) string {
	var repeated strings.Builder

	for range 5 {
		repeated.WriteString(character)
	}
	return repeated.String()
}

// Repeat6 Change the test so a caller can specify how many times the character is repeated and then fix the code
func Repeat6(character string, repeatTimes int) string {
	var repeated strings.Builder

	for range repeatTimes {
		repeated.WriteString(character)
	}
	return repeated.String()
}
