package edw

import (
	"errors"
)

// Trie Support restful.
// ':[variable_name]' Wildcard.

var ErrPatternFormatError = errors.New("current level exist Wildcard Node, no longer insert Node in this level")

type Node struct {
	Pattern  string
	part     string
	Wildcard bool
	children []*Node
	method   Method
	Handler  HandlerFunc
}

func newTrieRoot(rootHandler HandlerFunc) *Node {
	return &Node{
		Pattern:  "/",
		part:     "",
		Wildcard: false,
		children: nil,
		method:   REQUEST,
		Handler:  rootHandler,
	}
}

func (n *Node) insert(method Method, pattern string, parts []string, handler HandlerFunc, level int) {
	if len(parts) == level {
		n.Pattern = pattern
		n.method = method
		n.Handler = handler
		return
	}

	curPart := parts[level]
	child := n.insertMatchNode(curPart)

	if child == nil {
		child = &Node{
			part: curPart,
			// If current Node or parent Node is Wildcard,
			// current Node will be Wildcard.
			Wildcard: n.Wildcard || curPart[0] == ':' || curPart[0] == '*',
		}
		n.children = append(n.children, child)
	}

	child.insert(method, pattern, parts, handler, level+1)

	return
}

func (n *Node) insertMatchNode(part string) *Node {
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	return nil
}

func (n *Node) search(method Method, pattern string, parts []string, level int) *Node {
	if len(parts) == level {
		if n.suit(method) {
			return n
		}
		return nil
	}

	curPart := parts[level]

	if child := n.searchMatchNode(curPart); child != nil {
		return child.search(method, pattern, parts, level+1)
	}

	return nil
}

func (n *Node) searchMatchNode(part string) *Node {
	// exact match
	for _, child := range n.children {
		if child.part == part {
			return child
		}
	}
	// fuzzy match
	for _, child := range n.children {
		if child.Wildcard {
			return child
		}
	}
	return nil
}

// formatPattern
func formatUrlPath(pattern string) (string, error) {
	if pattern[0] != '/' {
		return "", ErrPatternFormatError
	}
	var ret []byte
	for i := range pattern {
		b := pattern[i]
		if b == '/' && len(ret) > 0 && ret[len(ret)-1] == '/' {
			return "", ErrPatternFormatError
		}
		if b == '?' {
			return "", ErrPatternFormatError
		}
		ret = append(ret, b)
	}

	return string(ret), nil
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

func (n *Node) suit(method Method) bool {
	if n.method == REQUEST {
		return true
	}
	return n.method == method
}
