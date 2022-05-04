package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name *string `json:"name,omitempty"`
	Age  *int    `json:"age,omitempty"`
}

func main() {
	bytes1, err := ioutil.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}

	persons := []Person{}
	json.Unmarshal([]byte(string(bytes1)), &persons)
	for _, person := range persons {
		fmt.Println("-----")
		if person.Name != nil {
			fmt.Println(*person.Name)
		}
		if person.Age != nil {
			fmt.Println(*person.Age)
		}
	}

	bytes2, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes2))
}
