package regex_tree

import (
	"regexp"
	"regexp/syntax"
	"strings"
)

type regex struct {
	static  string
	dynamic *regexp.Regexp
}

func newRegex(path string) (regex, error) {
	if re, err := syntax.Parse(path, syntax.Perl); err != nil {
		return regex{}, err
	} else if re.Op == syntax.OpLiteral {
		if reg, err := regexp.Compile(path); err != nil {
			return regex{}, err
		} else {
			literal, _ := reg.LiteralPrefix()
			return regex{static: literal}, nil
		}
	}
	reg, err := regexp.Compile("^" + path)
	if err != nil {
		return regex{}, err
	}
	return regex{dynamic: reg}, nil
}

func (r regex) commonPrefix(path string) (string, string, string, error) {
	if len(r.static) > 0 {
		return regexpCommonPrefix(regexp.QuoteMeta(r.static), path)
	}
	return regexpCommonPrefix(r.dynamic.String()[1:], path)
}

func (r regex) String() string {
	var fields []string
	if r.static != "" {
		fields = append(fields, "static: "+r.static)
	}
	if r.dynamic != nil {
		fields = append(fields, "dynamic: "+r.dynamic.String())
	}
	return "{ " + strings.Join(fields, ", ") + " }"
}
