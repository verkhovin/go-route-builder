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
		name: "/",
		methods: map[string]http.Handler{},
	}
	return &RouteBuilder{RoutingTree{&coreNode}}
}

type RoutingTree struct {
	core *RoutingTreeNode
}

type RoutingTreeNode struct {
	children map[string]*RoutingTreeNode
	name string
	methods map[string]http.Handler
}

func (t RoutingTree) add(pattern string, method string, handler http.Handler) {
	path := strings.Split(pattern, "/")
	node := t.core
	for _, value := range path {
		previousNode := node
		node = node.children[value]
		if node == nil {
			createChild(previousNode, value)
		} else {
			
		}
	}
}