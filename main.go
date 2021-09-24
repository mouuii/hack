package main

import (
	"fmt"
	"sort"
)

func main() {
	sl := sort.IntSlice([]int{89, 9, 81, 0, 2})
	fmt.Println(sl)
	sort.Sort(sl)
	fmt.Println(sl)
}
