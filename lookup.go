package regex_tree

import (
	"strings"
)

func (n *Node) Lookup(path string) (interface{}, []string) {
	var matched string
	var params []string
	if len(n.static) > 0 {
		if strings.HasPrefix(path, n.static) {
			matched = n.static
		}
	} else if params = n.dynamic.FindStringSubmatch(path); len(params) > 0 {
		// if no static, dynamic must be non nil, so need not to check dynamic is not nil.
		matched, params = params[0], params[1:]
	}
	if len(matched) == 0 {
		return nil, nil
	}
	if path = path[len(matched):]; len(path) == 0 {
		return n.data, params
	}
	for _, child := range n.children {
		if data, childParams := child.Lookup(path); data != nil {
			return data, append(params, childParams...)
		}
	}
	return nil, nil
}
