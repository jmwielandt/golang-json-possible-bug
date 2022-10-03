package main

import (
	"encoding/json"
	"fmt"
)

type Inner struct {
	FieldA string `json:"a"`
	FieldB int    `json:"b"`
}

type Out struct {
	FieldA string `json:"a"`
	FieldB int    `json:"b"`
	Inner  *Inner `json:"inner"` // pointer to inner
	X      []int  `json:"x"`
}

func main() {

	f1 := `{
		"a": "hi",
		"b": 4,
		"inner": {
			"a": "test 1",
			"b": 9
		},
		"x": [1, 2, 3]
	}`
	f2 := `{
		"a": "john",
		"b": 25,
		"inner": {
			"a": "smith",
			"b": 34
		},
		"x": [4, 5, 6]
	}`

	fs := []string{f1, f2}
	outs := make([]Out, 0)

	var x Out // defined outside the for loop
	for _, f := range fs {
		json.Unmarshal([]byte(f), &x)
		outs = append(outs, x)
	}

	for _, out := range outs {
		fmt.Printf("%#v\n", out)
		x, _ := json.Marshal(out)
		fmt.Println(string(x))
		fmt.Println()
	}
}
