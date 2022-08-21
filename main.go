package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(array1)
	array2 := lo.FlatMap(array1, func(element int, _ int) []int {
		if element%2 == 0 {
			return []int{element, element}
		}
		return []int{element, element, element}
	})
	fmt.Println(array2)
}
