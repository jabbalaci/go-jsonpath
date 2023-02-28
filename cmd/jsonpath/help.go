package main

import (
	"fmt"
	"strings"
)

func print_help() {
	text := `
jsonpath v{ver}
https://github.com/jabbalaci/go-jsonpath

Print every path / value pair in a JSON document.

Usage:

    jsonpath <input.json>
`
	text = strings.TrimLeft(text, "\n")
	text = strings.Replace(text, "{ver}", VERSION, 1)

	fmt.Println(text)
}
