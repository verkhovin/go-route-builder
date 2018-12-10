package routebuilder

import (
	"context"
	"net/http"
	"strings"
)

type RoutingTree struct {
	core *RoutingTreeNode
}

type RoutingTreeNode struct {
	children map[string] *RoutingTreeNode
	parametrizedChildren []*RoutingTreeNode
	name     string
	methods  map[string]http.Handler
}

func (t RoutingTree) add(pattern string, method string, handler http.Handler) {
	path := splitPath(pattern)
	node := t.core
	for _, value := range path {
		var child *RoutingTreeNode
		if isParam(value) {
			var ok bool
			if child, ok = node.getParametrizedChild(value); !ok {
				child = newNode(value)
				node.parametrizedChildren = append(node.parametrizedChildren, child)
			}
		} else {
			child = node.children[value]
			if child == nil {
				child = newNode(value)
				node.children[value] = child
			}
		}
		node = child
	}
	node.methods[method] = handler
}

func (node RoutingTreeNode) getParametrizedChild(paramId string) (*RoutingTreeNode, bool) {
	for _, v := range node.parametrizedChildren {
		if v.name == paramId {
			return v, true
		}
	}
	return nil, false
}
func isParam(s string) bool {
	return strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}")
}

func (t RoutingTree) getHandlerByUrl(url string, r *http.Request) http.Handler {
	path := splitPath(url)
	node := t.core
	method := r.Method
	for idx, value := range path {
		child := node.children[value]
		if child == nil {
			if paramChild, ok := node.tryAsParam(path[idx:], method); ok {
				ctx := context.WithValue(r.Context(), value, paramChild.name)
				r = r.WithContext(ctx) ///!!!!!
				child = paramChild
			} else {
				return notFoundHandler()
			}
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

func (node *RoutingTreeNode) tryAsParam(pathSlice []string, method string) (*RoutingTreeNode, bool) {
	paramChildren := node.parametrizedChildren
	if len(paramChildren) == 0 {
		return nil, false
	}

	firstChildNodeWithParamsIdx := -1
	for idx, child := range paramChildren {
		if len(pathSlice) == 1 {
			if _, ok := child.methods[method]; ok {
				return child, true
			}
		} else {
			if _, ok := child.children[pathSlice[1]]; ok {
				return child, true
			}
		}

		if firstChildNodeWithParamsIdx < 0 && len(child.parametrizedChildren) > 0 {
			firstChildNodeWithParamsIdx = idx
		}
	}

	if(firstChildNodeWithParamsIdx > 0) {
		return paramChildren[firstChildNodeWithParamsIdx], true
	}

	return nil, false
}

func (t RoutingTree) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler := t.getHandlerByUrl(request.URL.String(), request)
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
