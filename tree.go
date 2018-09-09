package regex_tree

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrorNoCommonPrefix = errors.New("no common prefix")
	ErrorAreadyExists   = errors.New("path aready exists")
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

func (n *Node) Add(path string, data interface{}) error {
	common, splitChildPath, childPath, err := n.commonPrefix(path)
	if err != nil {
		return err
	}
	if len(common) == 0 {
		return ErrorNoCommonPrefix
	}
	if len(splitChildPath) > 0 {
		if err := n.split(common, splitChildPath); err != nil {
			return err
		}
	}
	if len(childPath) == 0 {
		if n.data == nil {
			n.data = data
			return nil
		} else {
			return ErrorAreadyExists
		}
	}
	return n.addToChildren(childPath, data)
}

func (n *Node) addToChildren(path string, data interface{}) error {
	for _, child := range n.children {
		if err := child.Add(path, data); err != ErrorNoCommonPrefix {
			return err
		}
	}
	child, err := New(path, data)
	if err != nil {
		return err
	}
	// 静态路径优先匹配，所以将静态子节点放在动态子节点前边
	if l := len(n.children); l > 0 && len(child.static) > 0 && n.children[l-1].dynamic != nil {
		i := 0
		for ; i < l && len(n.children[i].static) > 0; i++ {
		}
		children := append(make([]*Node, 0, l+1), n.children[:i]...)
		children = append(children, child)
		n.children = append(children, n.children[i:]...)
	} else {
		n.children = append(n.children, child)
	}
	return nil
}

// 分裂为父节点和子节点
func (n *Node) split(path, childPath string) error {
	child, err := New(childPath, n.data)
	if err != nil {
		return err
	}
	child.children = n.children

	reg, err := newRegex(path)
	if err != nil {
		return err
	}
	n.regex = reg
	n.data = nil
	n.children = []*Node{child}
	return nil
}

func (n *Node) Lookup(path string) (interface{}, []string) {
	_, data, params := n.lookup(path)
	return data, params
}

func (n *Node) lookup(path string) (bool, interface{}, []string) {
	matched, params := n.match(path)
	if len(matched) == 0 {
		return false, nil, nil
	}

	childPath := path[len(matched):]
	if len(childPath) == 0 {
		return true, n.data, params
	}
	for _, child := range n.children {
		if ok, data, childParams := child.lookup(childPath); ok {
			if len(childParams) > 0 {
				params = append(params, childParams...)
			}
			return true, data, params
		}
	}
	return true, nil, nil
}

func (n *Node) String() string {
	return n.string("")
}

func (n *Node) string(indent string) string {
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
			children.WriteString(child.string(indent+"  ") + "\n")
		}
		fields = append(fields, fmt.Sprintf("children: [\n%s%s]", children.String(), indent))
	}

	return indent + "{ " + strings.Join(fields, ", ") + " }"
}
