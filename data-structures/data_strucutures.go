package datastructures

import (
	"errors"
	"fmt"
)

func LinearSearch() (bool, error) {
	fmt.Println("LINEAR SEARCH")

	key := 3
	datalist := []int{1, 2, 3, 4, 5, 6, 7, 100}

	for _, i := range datalist {
		defer fmt.Printf("%v", i)
		if i == key {
			return true, nil
		}
	}

	return false, errors.New("Did not find the entry")

}
