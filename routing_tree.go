package routebuilder

import (
	"net/http"
	"strings"
)

type RoutingTree struct {
	core *RoutingTreeNode
}

type RoutingTreeNode struct {
	children map[string]*RoutingTreeNode
	name     string
	methods  map[string]http.Handler
}

func (t RoutingTree) add(pattern string, method string, handler http.Handler) {
	path := splitPath(pattern)
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

func (t RoutingTree) getHandlerByUrl(url string, method string) http.Handler {
	path := splitPath(url)
	node := t.core
	for _, value := range path {
		child := node.children[value]
		if child == nil {
			return notFoundHandler()
		}
		node = child
	}
	handler, ok := node.methods[method]
	if !ok {
		return notFoundHandler()
	} else {
		return handler
	}
}

func (t RoutingTree) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler := t.getHandlerByUrl(request.URL.String(), request.Method)
	handler.ServeHTTP(writer, request)
}

func newNode(name string) *RoutingTreeNode {
	return &RoutingTreeNode{name: name,
		children: map[string]*RoutingTreeNode{},
		methods:  map[string]http.Handler{},
	}
}

func splitPath(path string) []string {
	splitFunc := func(char rune) bool{
		return char == '/'
	}
	return strings.FieldsFunc(path, splitFunc)
}
