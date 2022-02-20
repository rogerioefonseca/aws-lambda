// main.go
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"net/http"
	datastructures "rogerioefonseca/serverless/data-structures"
	defer_test "rogerioefonseca/serverless/defer"
)

func hello(event InputEvent) (string, error) {
	return "It works", nil
}

func dataStructures() {
	fmt.Println("##############")
	result, err := datastructures.LinearSearch()

	if err != nil {
		fmt.Println("entry not found")
		os.Exit(1)
	}

	fmt.Printf("%t", result)
}

func main() {
	defer_test.Foo()
	dataStructures()

	fmt.Println("")
	fmt.Println("##############")

	//keyboard.GetCity()
	//lambda.Stdefer_test(hello)
}

type InputEvent struct {
	Link string `json:"link"`
	Key  string `json:"key"`
}

func getImage(url string) (bytes []byte) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	bytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return bytes

}
