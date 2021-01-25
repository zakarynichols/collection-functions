package main

import (
	utils "collection-functions/utils"
	"fmt"
	"strings"
)

func main() {
	var strs = []string{"peach", "apple", "pear", "pineapple"}

	// 2
	fmt.Println(utils.FindIndex(strs, "pear"))

	// true
	fmt.Println(utils.Includes(strs, "pear"))

	// false
	fmt.Println(utils.Includes(strs, "grape"))

	// false
	fmt.Println(utils.Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "z")
	}))

	// false
	fmt.Println(utils.All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	// [apple pineapple]
	fmt.Println(utils.Filter(strs, func(v string) bool {
		return strings.Contains(v, "l")
	}))

	// [PEACH APPLE PEAR PINEAPPLE]
	fmt.Println(utils.Map(strs, strings.ToUpper))
}
