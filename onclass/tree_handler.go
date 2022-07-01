package main

import "strings"

type HandlerBasedOnTree struct {
	root *node
}

type node struct {
	path     string
	children []*node

	// 如果这是叶子节点，
	// 那么匹配上之后就可以调用该方法
	handler handlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	//TODO implement me
	panic("implement me")
}

func (h *HandlerBasedOnTree) RouteV4(method string, pattern string, handlerFunc handlerFunc) {
	// 将pattern按照URL的分隔符切割
	// 例如，/user/friends 将变成 [user, friends]
	// 将前后的/去掉，统一格式
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")
	// 当前指向根节点
	cur := h.root
	for index, path := range paths {
		// 从子节点里边找一个匹配到了当前 path 的节点
		matchChild, found := h.findMatchChild(cur, path)
		if found {
			cur = matchChild
		} else {
			// 为当前节点根据
			h.createSubTree(cur, paths[index:], handlerFunc)
			return
		}
	}
	// 离开了循环，说明我们加入的是短路径，
	// 比如说我们先加入了 /order/detail
	// 再加入/order，那么会走到这里
	cur.handler = handlerFunc
}

// findMatchChild 这里应该使用 *node，这样就不需要传参 root
// 因为节点自身需要寻找子节点，而不是使用节点的人知道怎么找
func (h *HandlerBasedOnTree) findMatchChild(root *node, path string) (*node, bool) {
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

// createSubTree
// handlerFn：命中之后叶子节点的业务处理逻辑
func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handlerFn handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		// user.children = [profile, home, friends]
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFn
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}
