package main

import (
	"fmt"

	"github.com/golang-collections/collections/stack"
)

func main() {
	st := stack.New()
	st.Push(10)
	st.Push(10)
	st.Push(10)
	st.Push(10)
	st.Push(10)
	st.Push(10)

	for i := 0; i < 10; i++ {
		fmt.Println(st.Pop(), st.Len())
	}
}
