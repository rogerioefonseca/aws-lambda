// main.go
package main

import (
	"fmt"
	"io/ioutil"

	"net/http"
	art "rogerioefonseca/serverless/defer"
	keyboard "rogerioefonseca/serverless/keyboard_input"
)

func hello(event InputEvent) (string, error) {
	return "It works", nil
}

func main() {
	fmt.Println(art.Test)
	keyboard.GetCity()
	//lambda.Start(hello)
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
