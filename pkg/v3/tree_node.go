package webv3

import (
	"strings"
)

const (

	// 根节点，只有根用这个
	nodeTypeRoot = iota

	// *
	nodeTypeAny

	// 路径参数
	nodeTypeParam

	// 正则
	nodeTypeReg

	// 静态，即完全匹配
	nodeTypeStatic // 最高优先级放到了最后
)

const any = "*"

// matchFunc 承担两个职责，一个是判断是否匹配，一个是在匹配之后
// 将必要的数据写入到 Context
// 所谓必要的数据，这里基本上是指路径参数
// v1：==
// v2：child.path = *
// child.path 是路径参数，路径参数写入到 context 里面去
// 正则：chiild.path.match(reg)
// v3：抽象 matchFunc
type matchFunc func(path string, c *Context) bool

//// Node 将 Node 抽象为接口
//type Node interface {
//	Match() bool
//	FindChild(string)
//}

type node struct {

	// /user/*id 输出命中的模式 /user/:id
	children []*node

	// a -> ab, ac
	//children2 map[byte][]*node

	//// children 是有序的，查找的时候可以二分查找，但是对任意匹配的需要特殊处理
	//children3 []*node

	//// a -> 97 -> 0 -> []*node，二维数组，字符 a 对应 ASCII 为 97，把 97 对应到下标为 0 的数组，比 map 高效
	//children4 [][]*node

	// 如果这是叶子节点，
	// 那么匹配上之后就可以调用该方法
	handler   handlerFunc
	matchFunc matchFunc

	// 原始的 pattern。注意，它不是完整的pattern，
	// 而是匹配到这个节点的pattern
	pattern  string
	nodeType int
}

// 静态节点
func newStaticNode(path string) *node {
	return &node{
		children: make([]*node, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			return path == p && p != "*"
		},
		nodeType: nodeTypeStatic,
		pattern:  path,
	}
}

func newRootNode(method string) *node {
	return &node{
		children: make([]*node, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			panic("never call me")
		},
		nodeType: nodeTypeRoot,
		pattern:  method,
	}
}

func newNode(path string) *node {
	if path == "*" {
		return newAnyNode()
	}
	if strings.HasPrefix(path, ":") {
		return newParamNode(path)
	}
	return newStaticNode(path)
}

// 通配符 * 节点
func newAnyNode() *node {
	return &node{
		// 因为我们不允许 * 后面还有节点，所以这里可以不用初始化
		//children: make([]*node, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			return true
		},
		nodeType: nodeTypeAny,
		pattern:  any,
	}
}

// 路径参数节点
func newParamNode(path string) *node {
	paramName := path[1:]
	return &node{
		children: make([]*node, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			if c != nil {
				c.PathParams[paramName] = p
			}
			// 如果自身是一个参数路由，
			// 然后又来一个通配符，我们认为是不匹配的
			return p != any
		},
		nodeType: nodeTypeParam,
		pattern:  path,
	}
}

// 正则节点
//func newRegNode(path string) *node {
//	// 依据你的规则拿到正则表达式
//  拿到正则表达式 p
//	return &node{
//		children: make([]*node, 0, 2),
//		matchFunc: func(p string, c *Context) bool {
//			// 怎么写？
//			// 正则匹配一下 p
//		},
//		nodeType: nodeTypeParam,
//		pattern: path,
//	}
//}

// 允许用户定义自己的节点类型

// 难点在于 newNode 的时候不知道使用哪种 nodeType
//
//type Factory func() *node
//
//var factories = map[int]Factory{}
//
//func RegisterFactory(t int, factory Factory) {
//	factories[t] = factory
//}

// 只注册一个 Factory

type Factory func(path string) *node

var factory Factory

func RegisterFactory(f Factory) {
	factory = f
}

func main() {
	RegisterFactory(func(path string) *node {
		// 是我自定义格式的路由
		if strings.HasPrefix(path, ":daming") {
			return &node{}
		} else {
			return newNode(path)
		}
	})
}
