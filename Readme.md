# regex\_tree
a regular expression prefix tree.

[![Build Status](https://travis-ci.org/lovego/regex_tree.svg?branch=master)](https://travis-ci.org/lovego/regex_tree)
[![Coverage Status](https://coveralls.io/repos/github/lovego/regex_tree/badge.svg?branch=master)](https://coveralls.io/github/lovego/regex_tree?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/regex_tree)](https://goreportcard.com/report/github.com/lovego/regex_tree)
[![GoDoc](https://godoc.org/github.com/lovego/regex_tree?status.svg)](https://godoc.org/github.com/lovego/regex_tree)

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
