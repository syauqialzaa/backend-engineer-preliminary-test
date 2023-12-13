package basic_concepts

import "fmt"

// class
type LazyClass struct {
	data int
}

func (lc *LazyClass) UnusedMethod() {
	fmt.Println("UnusedMethod called.")
}

func Coding() {
	lazyObj := LazyClass{ data: 42 }
	lazyObj.UnusedMethod()

	fmt.Println("Data:", lazyObj.data)
}