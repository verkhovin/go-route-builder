package routebuilder

import (
	"net/http"
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

func (rb RouteBuilder) Get(path string, handler http.Handler) {
	rb.routingTree.add(path, "GET", handler)
}

func (rb RouteBuilder) Post(path string, handler http.Handler) {
	rb.routingTree.add(path, "POST", handler)
}

func (rb RouteBuilder) Put(path string, handler http.Handler) {
	rb.routingTree.add(path, "Put", handler)
}

func (rb RouteBuilder) Delete(path string, handler http.Handler) {
	rb.routingTree.add(path, "Delete", handler)
}

func (rb RouteBuilder) Build() http.Handler {
	return http.Handler(rb.routingTree)
}

