package hyperscript

import (
	"bytes"
	"fmt"
	"runtime"
	"sort"
	"strings"
)

// Node represents a virtual dom node.
type Node struct {
	tag        string
	text       string
	raw        string
	attributes map[string]string
	classNames map[string]bool
	children   []Node
	err        error
}

// Attr specifies DOM attributes of a node.
type Attr map[string]interface{}

// Style provides a conenient way to specify the inline style a node.
type Style map[string]interface{}

// DangerousRaw outputs unescaped string. Use at your own risk.
type DangerousRaw string

func sortedKeyForInterfaceMap(m map[string]interface{}) []string {
	l := make([]string, len(m))
	i := 0
	for k := range m {
		l[i] = k
		i++
	}
	sort.Strings(l)
	return l
}

func parseTag(tag string) (string, string, []string, error) {
	tagEnd := -1
	idStart := -1
	id := ""
	classStart := -1
	var classNames []string

	tlen := len(tag)
	for i, c := range tag {
		if c == '>' || c == '<' {
			return "", "", nil, fmt.Errorf("special charactor not supported for tag/class/id: %s", tag)
		}
		if c == '#' {
			if tagEnd == -1 {
				tagEnd = i
			}
			if idStart >= 0 {
				return "", "", nil, fmt.Errorf("id specified more than once: %s", tag)
			}
			if classStart >= 0 {
				classNames = append(classNames, tag[classStart:i])
				classStart = -1
			}
			idStart = i + 1
			continue
		}
		if c == '.' {
			if tagEnd == -1 {
				tagEnd = i
			}
			if idStart >= 0 && id == "" {
				id = tag[idStart:i]
			}
			if classStart >= 0 {
				classNames = append(classNames, tag[classStart:i])
			}
			classStart = i + 1
		}
	}
	if classStart >= 0 {
		classNames = append(classNames, tag[classStart:tlen])
	} else if idStart >= 0 && id == "" {
		id = tag[idStart:tlen]
	} else {
		tagEnd = tlen
	}

	pureTag := "div"
	if tagEnd > 0 {
		pureTag = tag[:tagEnd]
	}
	return pureTag, id, classNames, nil
}

func parseAttrValue(key string, value interface{}) (string, error) {
	switch v := value.(type) {
	case string:
		return v, nil
	}
	return "", fmt.Errorf("invalid attribute value")
}

func parseClassNames(value interface{}, classNames map[string]bool) error {
	switch v := value.(type) {
	case string:
		for _, name := range strings.Split(v, " ") {
			classNames[name] = true
		}
		return nil
	case map[string]bool:
		for name, on := range v {
			classNames[name] = on
		}
		return nil
	default:
		return fmt.Errorf("invalid class attribute value %v", value)
	}
}

func errorfDepth(depth int, format string, a ...interface{}) error {
	_, fn, line, _ := runtime.Caller(depth + 1)
	msg := fmt.Sprintf(format, a...)
	return fmt.Errorf("%v, %v: %s", fn, line, msg)
}

// Add appends a child.
func (node *Node) Add(child *Node) {
	if child == nil {
		return
	}
	node.children = append(node.children, *child)
}

// H constructs a virtual DOM node.
func H(nodes ...interface{}) *Node {
	node := &Node{}
	if len(nodes) == 0 {
		return node
	}

	if tag, ok := nodes[0].(string); ok {
		tag, id, classNames, err := parseTag(tag)
		node.tag = tag
		if err != nil {
			node.err = err
			return node
		}
		node.attributes = map[string]string{}
		if id != "" {
			node.attributes["id"] = id
		}
		node.classNames = map[string]bool{}
		for i := range classNames {
			node.classNames[classNames[i]] = true
		}
		nodes = nodes[1:]
	}

	for _, n := range nodes {
		if n == nil {
			continue
		}
		switch cnode := n.(type) {
		case string:
			child := Node{}
			child.text = cnode
			node.children = append(node.children, child)
		case *Node:
			if cnode == nil {
				continue
			}
			if cnode.err != nil {
				node.err = cnode.err
				return node
			}
			node.children = append(node.children, *cnode)
		case Attr:
			for key, value := range cnode {
				if key == "class" {
					err := parseClassNames(value, node.classNames)
					if err != nil {
						node.err = err
					}
					continue
				}
				parsedValue, err := parseAttrValue(key, value)
				if err != nil {
					node.err = err
					return node
				}
				node.attributes[key] = parsedValue
			}
		case Style:
			if len(cnode) == 0 {
				continue
			}
			keys := sortedKeyForInterfaceMap(cnode)
			buf := bytes.Buffer{}
			for _, key := range keys {
				valueRaw := cnode[key]
				switch value := valueRaw.(type) {
				case string:
					fmt.Fprintf(&buf, "%s:%s;", key, value)
				default:
					node.err = errorfDepth(1, "invalid style value")
					return node
				}
			}
			node.attributes["style"] = buf.String()
		case DangerousRaw:
			child := Node{}
			child.raw = string(cnode)
			node.children = append(node.children, child)
		default:
			node.err = errorfDepth(1, "invalid argument %+v", n)
			return node
		}
	}
	return node
}
