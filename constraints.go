// constraints refer to a number of different types of
// limitations or requirements that must be satisfied

package main

import "fmt"

type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	fmt.Println("5 + 9 is:", getSumGen(5, 6))
	fmt.Println("3.7 + 6.76 is:", getSumGen(3.7, 6.76))
}
