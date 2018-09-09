package regex_tree

import (
	"errors"
)

var (
	ErrorNoCommonPrefix = errors.New("no common prefix")
	ErrorAlreadyExists  = errors.New("path already exists")
)

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
		}
		return ErrorAlreadyExists
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
