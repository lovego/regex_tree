package regex_tree

import (
	"bytes"
	"fmt"
	"strings"
)

type Node struct {
	regex
	data     interface{}
	children []*Node
}

func New(path string, data interface{}) (*Node, error) {
	reg, err := newRegex(path)
	if err != nil {
		return nil, err
	}
	return &Node{regex: reg, data: data}, nil
}

func (n *Node) String() string {
	return n.StringIndent("")
}

func (n *Node) StringIndent(indent string) string {
	var fields []string
	if n.static != "" {
		fields = append(fields, "static: "+n.static)
	}
	if n.dynamic != nil {
		fields = append(fields, "dynamic: "+n.dynamic.String())
	}
	if n.data != nil {
		fields = append(fields, "data: "+fmt.Sprint(n.data))
	}
	if len(n.children) > 0 {
		var children bytes.Buffer
		for _, child := range n.children {
			children.WriteString(child.StringIndent(indent+"  ") + "\n")
		}
		fields = append(fields, fmt.Sprintf("children: [\n%s%s]", children.String(), indent))
	}

	return indent + "{ " + strings.Join(fields, ", ") + " }"
}
