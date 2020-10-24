package regex_tree

import (
	"fmt"
)

func ExampleNode_Lookup() {
	root, _ := New("/", 0)
	root.Add("/users", 1)
	root.Add("/users/123/([0-9]+)", 22)
	root.Add("/users/([0-9]+)", 2)
	root.Add("/unix/([a-z]+)/([0-9.]+)", 3)
	root.Add("/users/root", 4)
	root.Add("/([0-9]+)", 5)
	fmt.Println(root.Lookup("/"))
	fmt.Println(root.Lookup("/users"))
	fmt.Println(root.Lookup("/users/123"))
	fmt.Println(root.Lookup("/users/123/456"))
	fmt.Println(root.Lookup("/unix/linux/4.4.0"))
	fmt.Println(root.Lookup("/users/root"))
	fmt.Println(root.Lookup("/987"))
	fmt.Println(root.Lookup("404"))
	fmt.Println(root.Lookup("/users404"))
	fmt.Println(root.Lookup("/unix/linux/4.4.0a"))

	// Output:
	// 0 []
	// 1 []
	// 2 [123]
	// 22 [456]
	// 3 [linux 4.4.0]
	// 4 []
	// 5 [987]
	// <nil> []
	// <nil> []
	// <nil> []
}

func ExampleNode_Lookup_loopback() {
	root, _ := New("/", 0)
	root.Add(`/(\w+)`, 1)
	root.Add(`/(\w+)/abc`, 2)
	root.Add(`/([a-z]+)/def`, 3)

	fmt.Println(root.Lookup("/"))
	fmt.Println(root.Lookup("/users"))
	fmt.Println(root.Lookup("/users/abc"))
	fmt.Println(root.Lookup("/users/def"))

	// Output:
	// 0 []
	// 1 [users]
	// 2 [users]
	// 3 [users]
}
