package regex_tree

import (
	"fmt"
)

func Example_stringCommonPrefix_empty1() {
	common, a, b := stringCommonPrefix("users", "managers")
	fmt.Printf("%s, %s, %s,", common, a, b)
	// Output: , users, managers,
}
func Example_stringCommonPrefix_same1() {
	common, a, b := stringCommonPrefix("/", "/")
	fmt.Printf("%s, %s, %s,", common, a, b)
	// Output: /, , ,
}

func Example_stringCommonPrefix_same2() {
	common, a, b := stringCommonPrefix("/users", "/users")
	fmt.Printf("%s, %s, %s,", common, a, b)
	// Output: /users, , ,
}

func Example_stringCommonPrefix_leftLonger1() {
	common, a, b := stringCommonPrefix("/users", "/")
	fmt.Printf("%s, %s, %s,", common, a, b)
	// Output: /, users, ,
}

func Example_stringCommonPrefix_leftLonger2() {
	common, a, b := stringCommonPrefix("/users/root", "/users")
	fmt.Printf("%s, %s, %s,", common, a, b)
	// Output: /users, /root, ,
}

func Example_stringCommonPrefix_rightLonger1() {
	common, a, b := stringCommonPrefix("/", "/users")
	fmt.Printf("%s, %s, %s,", common, a, b)
	// Output: /, , users,
}

func Example_stringCommonPrefix_rightLonger2() {
	common, a, b := stringCommonPrefix("/users", "/users/root")
	fmt.Printf("%s, %s, %s,", common, a, b)
	// Output: /users, , /root,
}

func Example_stringCommonPrefix_differentSuffix() {
	common, a, b := stringCommonPrefix("/users/list", "/users/root")
	fmt.Printf("%s, %s, %s,", common, a, b)
	// Output: /users/, list, root,
}

func Example_literalCommonPrefix_differentSuffix() {
	common, a, b, err := literalCommonPrefix(`/users\.html`, `/users\.htm`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /users\.htm, l, , <nil>
}

func Example_literalCommonPrefix_error1() {
	common, a, b, err := literalCommonPrefix(`(`, `/`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: , , , error parsing regexp: missing closing ): `(`
}

func Example_literalCommonPrefix_error2() {
	common, a, b, err := literalCommonPrefix(`/`, `)`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: , , , error parsing regexp: unexpected ): `)`
}

func Example_regexpCommonPrefix_empty1() {
	common, a, b, err := regexpCommonPrefix(`user_(\d+)/xyz`, `manager_(\d+)/def`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: , user_([0-9]+)/xyz, manager_([0-9]+)/def, <nil>
}

func Example_regexpCommonPrefix_empty2() {
	common, a, b, err := regexpCommonPrefix(`(\d+)/xyz`, `(\w+)/def`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: , ([0-9]+)/xyz, ([0-9A-Z_a-z]+)/def, <nil>
}

func Example_regexpCommonPrefix_empty3() {
	common, a, b, err := regexpCommonPrefix(`(\d+)`, `(\w+)`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: , ([0-9]+), ([0-9A-Z_a-z]+), <nil>
}

func Example_regexpCommonPrefix_same1() {
	common, a, b, err := regexpCommonPrefix(`/(\d+)`, `/([0-9]+)`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /([0-9]+), , , <nil>
}

func Example_regexpCommonPrefix_same2() {
	common, a, b, err := regexpCommonPrefix(`/user_(\d+)`, `/user_(\d+)`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /user_([0-9]+), , , <nil>
}

func Example_regexpCommonPrefix_leftLonger1() {
	common, a, b, err := regexpCommonPrefix(`/(\d+)`, `/`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /, ([0-9]+), , <nil>
}

func Example_regexpCommonPrefix_leftLonger2() {
	common, a, b, err := regexpCommonPrefix(`/users/(\d+)`, `/users/`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /users/, ([0-9]+), , <nil>
}

func Example_regexpCommonPrefix_leftLonger3() {
	common, a, b, err := regexpCommonPrefix(`/users/(\w+)/(\d+)`, `/users/(\w+)`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /users/([0-9A-Z_a-z]+), /([0-9]+), , <nil>
}

func Example_regexpCommonPrefix_rightLonger1() {
	common, a, b, err := regexpCommonPrefix(`/`, `/(\d+)`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /, , ([0-9]+), <nil>
}

func Example_regexpCommonPrefix_rightLonger2() {
	common, a, b, err := regexpCommonPrefix(`/users/`, `/users/(\d+)`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /users/, , ([0-9]+), <nil>
}

func Example_regexpCommonPrefix_differentSuffix1() {
	common, a, b, err := regexpCommonPrefix(`/user_([0-9]+)/xyz`, `/user_([0-9]+)/def`)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /user_([0-9]+)/, xyz, def, <nil>
}

func Example_regexpCommonPrefix_differentSuffix2() {
	common, a, b, err := regexpCommonPrefix("users/([0-9]+)", "users/managers")
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: users/, ([0-9]+), managers, <nil>
}

func Example_regexpCommonPrefix_differentSuffix3() {
	common, a, b, err := regexpCommonPrefix(
		"/([a-z]+)/members/([0-9]+)", "/([a-z]+)/managers/([0-9]+)",
	)
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: /([a-z]+)/m, embers/([0-9]+), anagers/([0-9]+), <nil>
}

func Example_regexpCommonPrefix_error1() {
	common, a, b, err := regexpCommonPrefix("(", "/")
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: , , , error parsing regexp: missing closing ): `(`
}

func Example_regexpCommonPrefix_error2() {
	common, a, b, err := regexpCommonPrefix("/", ")")
	fmt.Printf("%s, %s, %s, %v", common, a, b, err)
	// Output: , , , error parsing regexp: unexpected ): `)`
}
