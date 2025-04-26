package iteration

import "strings"

const repeatedTime = 5

// func Repeat(str string) string {

// 	var answer []string

// 	for range repeatedTime {
// 		answer = append(answer, str)
// 	}

// 	return strings.Join(answer, "")
// }

func Repeat(str string) string {
	var s strings.Builder

	for range repeatedTime {
		// can use WriteString
		s.WriteString(str) // 這個比下面快很多

		// fmt.Fprintf(&s, "%s", str) // 也可以用fmt 寫進writer, 這個最慢
	}

	return s.String()
}
