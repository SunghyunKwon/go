package main

import (
	"fmt"

	"./mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("Hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	definition, err := dictionary.Search("Hello")
	fmt.Println(definition)

}
