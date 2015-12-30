package hyperscript

import (
	"fmt"
	"strings"
)

// Node represents a virtual dom node.
type Node struct {
	tag        string
	text       string
	attributes map[string]string
	classNames map[string]bool
	children   []Node
	err        error
}

// Attr specifies DOM attributes of a node.
type Attr map[string]interface{}

// Style provides a conenient way to specify the inline style a node.
type Style map[string]interface{}

var byteSpace = []byte(" ")

func parseTag(tag string) (string, string, []string, error) {
	tagEnd := -1
	idStart := -1
	id := ""
	classStart := -1
	var classNames []string

	tlen := len(tag)
	for i, c := range tag {
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
	default:
		return fmt.Errorf("invalid class attribute value %v", value)
	}
}

// H constructs a virtual DOM node.
func H(tag string, nodes ...interface{}) *Node {
	tag, id, classNames, err := parseTag(tag)
	node := &Node{tag: tag}
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
	for _, n := range nodes {
		switch cnode := n.(type) {
		case string:
			child := Node{}
			child.text = cnode
			node.children = append(node.children, child)
		case *Node:
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
		default:
			node.err = fmt.Errorf("invalid argument %v", n)
			return node
		}
	}
	return node
}
