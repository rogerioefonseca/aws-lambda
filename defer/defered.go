package defer_test

import "fmt"

var Test = "Name Rogerio"

func Foo() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
