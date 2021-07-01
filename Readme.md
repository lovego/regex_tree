# regex\_tree
a regular expression prefix tree.

[![Build Status](https://github.com/lovego/regex_tree/actions/workflows/go.yml/badge.svg)](https://github.com/lovego/regex_tree/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/lovego/regex_tree/badge.svg?branch=master)](https://coveralls.io/github/lovego/regex_tree)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/regex_tree)](https://goreportcard.com/report/github.com/lovego/regex_tree)
[![Documentation](https://pkg.go.dev/badge/github.com/lovego/regex_tree)](https://pkg.go.dev/github.com/lovego/regex_tree@v0.0.1)


## Usage
```go
package main

import (
	"fmt"
	"github.com/lovego/regex_tree"
)

func main() {
	root, err := regex_tree.New("/", 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	root.Add("/users", 2)
	root.Add(`/users/(\d+)`, 3)

	fmt.Println(root.Lookup("/"))
	fmt.Println(root.Lookup("/users"))
	fmt.Println(root.Lookup("/users/1013"))
	fmt.Println(root.Lookup("/users/a013"))

	// Output:
	// 1 []
	// 2 []
	// 3 [1013]
	// <nil> []
}
```
