package regex_tree

import (
	"strings"
)

// if a child matches, now only this child is looked up, the other children will not be looked up.
func (n *Node) Lookup(path string) (interface{}, []string) {
	var matched string
	var params []string
	if len(n.static) > 0 {
		if strings.HasPrefix(path, n.static) {
			matched = n.static
		}
	} else if params = n.dynamic.FindStringSubmatch(path); len(params) > 0 {
		matched, params = params[0], params[1:]
	}
	if len(matched) == 0 {
		return nil, nil
	}
	var childParams []string
loop:
	for {
		if path = path[len(matched):]; len(path) == 0 {
			return n.data, params
		}
		for _, child := range n.children {
			if len(child.static) > 0 {
				if strings.HasPrefix(path, child.static) {
					matched = child.static
					n = child
					continue loop
				}
			} else if childParams = child.dynamic.FindStringSubmatch(path); len(childParams) > 0 {
				matched = childParams[0]
				if len(childParams) > 1 {
					params = append(params, childParams[1:]...)
				}
				n = child
				continue loop
			}
		}
		return nil, nil
	}
	return nil, nil
}
