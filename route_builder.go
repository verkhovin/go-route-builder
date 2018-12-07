package GoFluentRouter

import (
	"net/http"
	"strings"
)

type RouteBuilder struct {
	routingTree RoutingTree
}

func NewBuilder() *RouteBuilder {
	coreNode := RoutingTreeNode{
		children: map[string]*RoutingTreeNode{},
		name:     "/",
		methods:  map[string]http.Handler{},
	}
	return &RouteBuilder{RoutingTree{&coreNode}}
}

type RoutingTree struct {
	core *RoutingTreeNode
}

type RoutingTreeNode struct {
	children map[string]*RoutingTreeNode
	name     string
	methods  map[string]http.Handler
}

func (t RoutingTree) add(pattern string, method string, handler http.Handler) {
	path := strings.FieldsFunc(pattern, splitFunc)
	node := t.core
	for _, value := range path {
		child := node.children[value]
		if child == nil {
			child = newNode(value)
			node.children[value] = child
		}
		node = child
	}
	node.methods[method] = handler
}

func newNode(name string) *RoutingTreeNode {
	return &RoutingTreeNode{name: name,
		children: map[string]*RoutingTreeNode{},
		methods:  map[string]http.Handler{},
	}
}

func splitFunc(char rune) bool {
	return char == '/'
}
