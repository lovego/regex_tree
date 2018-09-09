package regex_tree

import (
	"fmt"
)

func Example_usage() {
	root, err := New("/", 1)
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

func ExampleNew() {
	fmt.Println(New("/", 1))
	fmt.Println(New("/users", 2))
	fmt.Println(New("/[a-z]+", 3))
	fmt.Println(New("/users/[0-9]+", 4))
	fmt.Println(New(`/users/\d+`, 5))
	fmt.Println(New(`(`, 6))
	// Output:
	// { static: /, data: 1 } <nil>
	// { static: /users, data: 2 } <nil>
	// { dynamic: ^/[a-z]+, data: 3 } <nil>
	// { dynamic: ^/users/[0-9]+, data: 4 } <nil>
	// { dynamic: ^/users/\d+, data: 5 } <nil>
	// <nil> error parsing regexp: missing closing ): `(`
}
