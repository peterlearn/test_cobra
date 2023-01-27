package main

import (
	"encoding/json"
	"fmt"
)

type test struct {
	A string
	B int
}

func main0003() {
	output, err := json.MarshalIndent(test{
		A: "a",
		B: 1,
	}, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(output))
}
