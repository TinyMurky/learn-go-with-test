package integer

import "fmt"

// Add takes two integers and returns the sum of them.
func Add(a, b int) int {
	return a + b
}

func main() {
	c := Add(1, 2)
	fmt.Printf("Answer is : %v", c)
}
