package main

import "fmt"

func main() {

	numberSum := []int{20, 22, 67}
	fmt.Println("Sum", Sum[int](numberSum))

	numberSub := []int{100, 50, 25}
	fmt.Println("Sub", Sub[int](numberSub))

}

// --------------------------------------------
// Declaring in func
func Sum[V int64 | float64 | int](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

// --------------------------------------------
// Declaring out func
type Interface interface {
	int64 | float64 | int
}

func Sub[V Interface](m []V) V {
	s := m[0]
	for _, v := range m {
		s -= v
	}

	return s
}
