package engine

import (
	"errors"
)

// Trie Support restful.
// ':[variable_name]' wildcard.

// ErrWildcardOnlyOne If curr child node is wildcard match node,
// must throw an exception,
// because a request only match a handler.
var ErrWildcardOnlyOne = errors.New("current level exist wildcard node, no longer insert node in this level")

type node struct {
	pattern  string
	part     string
	wildcard bool
	children []*node
}

func (n *node) insert(pattern string, parts []string, level int) (err error) {
	if len(parts) == level {
		n.pattern = pattern
		return
	}

	curPart := parts[level]
	child, err := n.insertMatchNode(curPart)
	if err != nil {
		return err
	}

	if child == nil {
		child = &node{
			part: curPart,
			// If current node or parent node is wildcard,
			// current node will be wildcard.
			wildcard: n.wildcard || curPart[0] == ':' || curPart[0] == '*',
		}
		// If a level exist a wildcard node that only one node.
		if child.wildcard && len(n.children) > 0 {
			return ErrWildcardOnlyOne
		}
		n.children = append(n.children, child)
	}

	if err = child.insert(pattern, parts, level+1); err != nil {
		return err
	}

	return
}

func (n *node) insertMatchNode(part string) (*node, error) {
	// If a level exist a wildcard node that only one node.
	if len(n.children) == 1 && n.children[0].wildcard {
		return nil, ErrWildcardOnlyOne
	}
	for _, child := range n.children {
		if child.part == part {
			return child, nil
		}
	}
	return nil, nil
}

func (n *node) search(pattern string, parts []string, level int) *node {
	if len(parts) == level {
		// TODO parse wildcard k-v.
		return n
	}

	curPart := parts[level]

	if child := n.searchMatchNode(curPart); child != nil {
		return child.search(pattern, parts, level+1)
	}

	return nil
}

func (n *node) searchMatchNode(part string) *node {
	for _, child := range n.children {
		if child.wildcard || child.part == part {
			return child
		}
	}
	return nil
}

func getParts(urlPath string) (parts []string) {
	var p []byte
	for i := range urlPath {
		b := urlPath[i]
		if b != '/' {
			p = append(p, b)
		} else if b == '/' && len(p) > 0 {
			parts = append(parts, string(p))
			p = p[:0]
		}
	}
	if len(p) > 0 {
		parts = append(parts, string(p))
	}
	return parts
}
