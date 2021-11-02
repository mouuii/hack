package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"
)

func intersect(s [][]string) []string {
	tmp := sets.String{}
	for i, v := range s {
		t := sets.NewString(v...)
		if i == 0 {
			tmp = t
		}
		tmp = t.Intersection(tmp)
	}
	return tmp.List()
}

func main() {
	s1 := []string{"ss", "x"}
	s11 := []string{"vvv", "x"}
	ssss := intersect([][]string{s1, s11})
	fmt.Println(ssss)

	s2 := []string{"v", "x"}
	s22 := []string{"vvv", "vv"}
	sssss := intersect([][]string{s2, s22})
	fmt.Println(sssss)
}
