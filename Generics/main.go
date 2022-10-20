package main

import "fmt"

func SmallestInteger[T interface{}](a T, b T) T {

	if a < b {
		return a
	}

	return b

}

type Pessoa struct {
	Name string
}

func main() {
	fmt.Println(IsEquals[string]("1", "2"))
	fmt.Println(IsEquals[int](1, 1))
	fmt.Println(IsEquals[bool](true, true))
	fmt.Println(IsEquals[Pessoa](Pessoa{"Tiago"}, Pessoa{"Tiago"}))
}
