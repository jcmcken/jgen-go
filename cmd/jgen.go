package main

import (
	"encoding/json"
	"fmt"
	"github.com/jcmcken/jgen-go/parser"
	"github.com/peterbourgon/mergemap"
	"os"
)

func main() {
	result := make(map[string]interface{})

	for _, arg := range os.Args[1:] {
		tree := parser.Parse(arg)
		result = mergemap.Merge(tree, result)
	}

	out, err := json.MarshalIndent(&result, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
