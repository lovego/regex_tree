package regex_tree

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func Example_newRegex_simple() {
	r, err := newRegex("/")
	fmt.Printf("%+v %v\n", r, err)
	r, err = newRegex(`/index\.html`)
	fmt.Printf("%+v %v\n", r, err)
	r, err = newRegex(`/index.html`)
	fmt.Printf("%+v %v\n", r, err)
	r, err = newRegex(`(`)
	fmt.Printf("%+v %v\n", r, err)
	// Output:
	// { static: / } <nil>
	// { static: /index.html } <nil>
	// { dynamic: ^/index.html } <nil>
	// {  } error parsing regexp: missing closing ): `(`
}

func Example_regex_commonPrefix_static() {
	r, _ := newRegex("/")
	common, a, b, err := r.commonPrefix("/")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex("users")
	common, a, b, err = r.commonPrefix("members")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex("/users")
	common, a, b, err = r.commonPrefix("/")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex("users")
	common, a, b, err = r.commonPrefix("([0-9]+)")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex("/")
	common, a, b, err = r.commonPrefix("/users")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex("users/managers")
	common, a, b, err = r.commonPrefix("users/([0-9]+)")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	// Output:
	// /, , , <nil>
	// , users, members, <nil>
	// /, users, , <nil>
	// , users, ([0-9]+), <nil>
	// /, , users, <nil>
	// users/, managers, ([0-9]+), <nil>
}

func Example_regex_commonPrefix_dynamic() {
	r, _ := newRegex("/([a-z]+)")
	common, a, b, err := r.commonPrefix("/")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex("users/([0-9]+)")
	common, a, b, err = r.commonPrefix("members")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex("/([a-z]+)")
	common, a, b, err = r.commonPrefix("/([a-z]+)/members")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex("users/([0-9]+)")
	common, a, b, err = r.commonPrefix("users/([a-z]+)")
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	r, _ = newRegex(`user\.s/([0-9]+)`)
	common, a, b, err = r.commonPrefix(`user\.s/managers`)
	fmt.Printf("%s, %s, %s, %v\n", common, a, b, err)
	// Output:
	// /, ([a-z]+), , <nil>
	// , users/([0-9]+), members, <nil>
	// /([a-z]+), , /members, <nil>
	// users/, ([0-9]+), ([a-z]+), <nil>
	// user\.s/, ([0-9]+), managers, <nil>
}

func BenchmarkStringHasPrefix(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i <= b.N; i++ {
		if !strings.HasPrefix("/company-skus/search", "/company-skus") {
			b.Error("not matched")
		}
	}
}

var testRegexp = regexp.MustCompile("^/company-skus")

func BenchmarkRegexpMatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i <= b.N; i++ {
		if !testRegexp.MatchString("/company-skus/search") {
			b.Error("not matched")
		}
	}
}

func BenchmarkRegexpFindSubmatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i <= b.N; i++ {
		if len(testRegexp.FindStringSubmatch("/company-skus/search")) == 0 {
			b.Error("not matched")
		}
	}
}
