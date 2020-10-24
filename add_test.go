package regex_tree

import (
	"fmt"
	"regexp"
)

func ExampleNode_Add_simple1() {
	root, _ := New("/", 0)
	fmt.Println(root.Add("(", 3))
	root.Add("/users", 1)
	root.Add(`/users/(\d+)`, 2)
	fmt.Println(root)
	// Output:
	// error parsing regexp: missing closing ): `(`
	// { static: /, data: 0, children: [
	//   { static: users, data: 1, children: [
	//     { dynamic: ^/([0-9]+), data: 2 }
	//   ] }
	// ] }
}

func ExampleNode_Add_simple2() {
	root, _ := New("/", 0)
	root.Add("/users", 1)
	root.Add("/users/([0-9]+)", 2)
	root.Add("/unix/([a-z]+)", 3)
	root.Add("/users/root", 4)
	root.Add("/([0-9]+)", 5)
	fmt.Println(root)
	// Output:
	// { static: /, data: 0, children: [
	//   { static: u, children: [
	//     { static: sers, data: 1, children: [
	//       { static: /, children: [
	//         { static: root, data: 4 }
	//         { dynamic: ^([0-9]+), data: 2 }
	//       ] }
	//     ] }
	//     { dynamic: ^nix/([a-z]+), data: 3 }
	//   ] }
	//   { dynamic: ^([0-9]+), data: 5 }
	// ] }
}

func ExampleNode_Add_conflict1() {
	root, _ := New("/", 0)
	fmt.Println(root.Add("/", 1))
	// Output: path already exists
}

func ExampleNode_Add_conflict2() {
	root, _ := New("/users", 0)
	fmt.Println(root.Add("/users", 1))
	// Output: path already exists
}

func ExampleNode_Add_conflict3() {
	root, _ := New("/users", 0)
	root.Add("/", 1)
	fmt.Println(root.Add("/users", 2))
	// Output: path already exists
}

func ExampleNode_Add_conflict4() {
	root, _ := New("/users/active", 0)
	root.Add("/", 1)
	root.Add("/users", 2)
	fmt.Println(root.Add("/users/active", 3))
	// Output: path already exists
}

func ExampleNode_Add_conflict5() {
	root, _ := New("/users/([0-9]+)", 0)
	root.Add("/", 1)
	root.Add("/users", 2)
	fmt.Println(root.Add("/users/([0-9]+)", 3))
	// Output: path already exists
}

func ExampleNode_addToChildren_static1() {
	n, _ := New("/", 0)
	fmt.Println(n.addToChildren(`(`, 1))
	n.addToChildren("users", 1)
	fmt.Println(n)
	fmt.Println(n.addToChildren(`(`, 1))
	// Output:
	// error parsing regexp: missing closing ): `(`
	// { static: /, data: 0, children: [
	//   { static: users, data: 1 }
	// ] }
	// error parsing regexp: missing closing ): `(`
}

func ExampleNode_addToChildren_static2() {
	n, _ := New("/u", 0)
	n.children = []*Node{
		{regex: regex{dynamic: regexp.MustCompile("^/")}},
	}
	n.addToChildren("sers", 1)
	fmt.Println(n)
	// Output:
	// { static: /u, data: 0, children: [
	//   { static: sers, data: 1 }
	//   { dynamic: ^/ }
	// ] }
}

func ExampleNode_addToChildren_static3() {
	n, _ := New("/u", 0)
	n.children = []*Node{
		{regex: regex{static: "nix"}},
		{regex: regex{dynamic: regexp.MustCompile("^/1")}},
		{regex: regex{dynamic: regexp.MustCompile("^/2")}},
	}
	n.addToChildren("sers", 1)
	fmt.Println(n)
	// Output:
	// { static: /u, data: 0, children: [
	//   { static: nix }
	//   { static: sers, data: 1 }
	//   { dynamic: ^/1 }
	//   { dynamic: ^/2 }
	// ] }
}

func ExampleNode_addToChildren_dynamic1() {
	n, _ := New("/u", 0)
	n.children = []*Node{
		{regex: regex{static: "sers"}},
		{regex: regex{dynamic: regexp.MustCompile("^/")}},
	}
	n.addToChildren("[0-9]+", 1)
	fmt.Println(n)
	// Output:
	// { static: /u, data: 0, children: [
	//   { static: sers }
	//   { dynamic: ^/ }
	//   { dynamic: ^[0-9]+, data: 1 }
	// ] }
}

func ExampleNode_split_static1() {
	n, _ := New("/users", 0)
	n.split("/", "users")
	fmt.Println(n)
	fmt.Println(n.split(`(`, "/"))
	fmt.Println(n.split(`/`, ")"))
	// Output:
	// { static: /, children: [
	//   { static: users, data: 0 }
	// ] }
	// error parsing regexp: missing closing ): `(`
	// error parsing regexp: unexpected ): `)`
}

func ExampleNode_split_static2() {
	n, _ := New("/users/managers", 0)
	n.split("/users/", "managers")
	fmt.Println(n)
	// Output:
	// { static: /users/, children: [
	//   { static: managers, data: 0 }
	// ] }
}

func ExampleNode_split_dynamic1() {
	n, _ := New("/[a-z]+", 0)
	n.split("/", "[a-z]+")
	fmt.Println(n)
	// Output:
	// { static: /, children: [
	//   { dynamic: ^[a-z]+, data: 0 }
	// ] }
}

func ExampleNode_split_dynamic2() {
	n, _ := New(`/users/[0-9]+`, 0)
	n.split("/u", "sers/[0-9]+")
	fmt.Println(n)
	// Output:
	// { static: /u, children: [
	//   { dynamic: ^sers/[0-9]+, data: 0 }
	// ] }
}

func ExampleNode_split_dynamic3() {
	n, _ := New(`/([a-z]+)/([0-9]+)`, 0)
	n.split("/([a-z]+)/", "([0-9]+)")
	fmt.Println(n)
	// Output:
	// { dynamic: ^/([a-z]+)/, children: [
	//   { dynamic: ^([0-9]+), data: 0 }
	// ] }
}

func ExampleNode_split_dynamic4() {
	n, _ := New("/users/[0-9]+", 0)
	n.split("/users/", "[0-9]+")
	fmt.Println(n)
	// Output:
	// { static: /users/, children: [
	//   { dynamic: ^[0-9]+, data: 0 }
	// ] }
}
