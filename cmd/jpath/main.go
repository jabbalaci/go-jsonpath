package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Traverse the object recursively and print every path / value pair.
func traverse(path string, obj interface{}) {
	switch obj.(type) {
	case []interface{}:
		for i, subnode := range obj.([]interface{}) {
			traverse(fmt.Sprintf("%s[%d]", path, i), subnode)
		}
	case map[string]interface{}:
		for k, v := range obj.(map[string]interface{}) {
			traverse(fmt.Sprintf("%s[%q]", path, k), v)
		}
	case string:
		fmt.Printf("%s => %q\n", path, obj)
	case float64:
		f := obj.(float64)
		num := int(f)
		if f == float64(num) {
			// the number is an integer
			fmt.Printf("%s => %d\n", path, num)
		} else {
			fmt.Printf("%s => %f\n", path, f)
		}
	default:
		fmt.Printf("%s => %v\n", path, obj)
	}
}

// Read the JSON file and return its content as a Go data structure.
func readFile(fpath string) (map[string]interface{}, error) {
	file, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data map[string]interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Process the given JSON file.
func process(fname string) error {
	data, err := readFile(fname)
	if err != nil {
		return err
	}
	traverse("root", data)
	return nil
}

func main() {
	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	if arg == "" || arg == "-h" || arg == "--help" {
		print_help()
		os.Exit(0)
	}

	err := process(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
