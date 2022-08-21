package main

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
)

func outputFlatMap() {
	array1 := []int{1, 2, 3, 4, 5, 6}
	array2 := lo.FlatMap(array1, func(element int, _ int) []int {
		if element%2 == 0 {
			return []int{element, element}
		}
		return []int{element, element, element}
	})
	fmt.Println("outputFlatMap: ", array2)
}

func outputMap() {
	samples := []string{"A", "B", "C"}
	indices := lo.Map(samples, func(sample string, idx int) int {
		return idx
	})
	fmt.Println("outputMap: ", indices)
}

func outputFilter() {
	samples := []string{"Apple", "Anchor", "Bee"}
	filtered := lo.Filter(samples, func(sample string, idx int) bool {
		return strings.HasPrefix(sample, "A")
	})
	fmt.Println("outputFilter: ", filtered)
}

func outputReduce() {
	samples := []int{2, 2, 2, 2, 2}
	reduced := lo.Reduce(samples, func(result int, sample int, idx int) int {
		return result * sample
	}, 1)
	fmt.Println("outputReduce: ", reduced)
}

func main() {
	outputFlatMap()
	outputMap()
	outputFilter()
	outputReduce()
}
