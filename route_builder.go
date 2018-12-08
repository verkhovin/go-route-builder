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

func (rb RouteBuilder) Get(path string, handler http.Handler) RouteBuilder{
	rb.routingTree.add(path, "GET", handler)
	return rb
}

func (rb RouteBuilder) Post(path string, handler http.Handler) RouteBuilder{
	rb.routingTree.add(path, "POST", handler)
	return rb
}

func (rb RouteBuilder) Put(path string, handler http.Handler) RouteBuilder{
	rb.routingTree.add(path, "PUT", handler)
	return rb
}

func (rb RouteBuilder) Delete(path string, handler http.Handler) RouteBuilder{
	rb.routingTree.add(path, "GET", handler)
	return rb
}

func (rb RouteBuilder) Build() http.Handler {
	return http.Handler(rb.routingTree)
}

