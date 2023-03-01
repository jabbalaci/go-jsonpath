package main

import (
	"fmt"
	"strings"
)

func print_help() {
	text := `
jpath {ver}
https://github.com/jabbalaci/go-jsonpath

Print every path / value pair in a JSON document.

Usage:

    jpath <input.json>
`
	text = strings.TrimLeft(text, "\n")
	text = strings.Replace(text, "{ver}", VERSION, 1)

	fmt.Println(text)
}
