package regex_tree

import (
	"regexp"
	"regexp/syntax"
)

// 正则表达式公共前缀
func regexpCommonPrefix(aStr, bStr string) (string, string, string, error) {
	a, err := syntax.Parse(aStr, syntax.Perl)
	if err != nil {
		return "", "", "", err
	}
	b, err := syntax.Parse(bStr, syntax.Perl)
	if err != nil {
		return "", "", "", err
	}
	if a.Equal(b) {
		return a.String(), "", "", nil
	}

	var aSlice, bSlice []*syntax.Regexp
	if a.Op == syntax.OpConcat {
		aSlice = a.Sub
	} else {
		aSlice = []*syntax.Regexp{a}
	}
	if b.Op == syntax.OpConcat {
		bSlice = b.Sub
	} else {
		bSlice = []*syntax.Regexp{b}
	}
	return regexpSliceCommonPrefix(aSlice, bSlice)
}

func regexpSliceCommonPrefix(aSlice, bSlice []*syntax.Regexp) (
	common string, a string, b string, err error,
) {
	size := len(aSlice)
	if size > len(bSlice) {
		size = len(bSlice)
	}
	i := 0
	for ; i < size; i++ {
		if aSlice[i].Equal(bSlice[i]) {
			common += aSlice[i].String()
		} else if aSlice[i].Op == syntax.OpLiteral && bSlice[i].Op == syntax.OpLiteral {
			if comm, aLiteral, bLiteral, err := literalCommonPrefix(
				aSlice[i].String(), bSlice[i].String(),
			); err != nil {
				return "", "", "", err
			} else {
				common += comm
				a, b = aLiteral, bLiteral
			}
			i++
			break
		} else {
			break
		}
	}
	for j := i; j < len(aSlice); j++ {
		a += aSlice[j].String()
	}
	for j := i; j < len(bSlice); j++ {
		b += bSlice[j].String()
	}
	return
}

func literalCommonPrefix(a, b string) (string, string, string, error) {
	aReg, err := regexp.Compile(a)
	if err != nil {
		return "", "", "", err
	}
	bReg, err := regexp.Compile(b)
	if err != nil {
		return "", "", "", err
	}
	aStr, _ := aReg.LiteralPrefix()
	bStr, _ := bReg.LiteralPrefix()
	common, a, b := stringCommonPrefix(aStr, bStr)
	return regexp.QuoteMeta(common), regexp.QuoteMeta(a), regexp.QuoteMeta(b), nil
}

// 字符串公共前缀
func stringCommonPrefix(a, b string) (string, string, string) {
	size := len(a)
	if size > len(b) {
		size = len(b)
	}
	for i := 0; i < size; i++ {
		if a[i] != b[i] {
			return a[:i], a[i:], b[i:]
		}
	}
	return a[:size], a[size:], b[size:]
}
