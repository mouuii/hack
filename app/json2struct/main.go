package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	person1 := person{"小明", 18}
	if result, err := json.Marshal(&person1); err == nil {
		fmt.Println(string(result))
	}
	bs, _ := json.Marshal(person1)
	txt := string(bs)
	fmt.Println(txt)
	txtRal := `{"Name":"小明","Age":18}`
	var person2 person
	err := json.Unmarshal([]byte(txtRal), &person2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person2)
}
