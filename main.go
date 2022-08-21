package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	names := lo.Uniq([]string{"Samuel", "Marc", "Samuel"})
	fmt.Println(names)
}
