package hyperscript

import (
	"bytes"
	"fmt"
	"html"
	"io"
	"strings"
)

type Node struct {
	tag        string
	text       string
	attributes map[string]string
	classNames map[string]bool
	children   []Node
	err        error
}

type Attr map[string]interface{}
type Style map[string]interface{}

var byteSpace = []byte(" ")

func (node *Node) WriteTo(w io.Writer, indent string, indentStep string) error {
	if node.err != nil {
		return node.err
	}
	if node.text != "" {
		w.Write([]byte(html.EscapeString(node.text)))
		return nil
	}
	w.Write([]byte(indent))
	w.Write([]byte("<"))
	w.Write([]byte(node.tag))

	if len(node.classNames) > 0 {
		w.Write([]byte(` class="`))
		sep := []byte{}
		for k, v := range node.classNames {
			if !v {
				continue
			}
			w.Write(sep)
			w.Write([]byte(k))
			sep = byteSpace
		}
		w.Write([]byte(`"`))
	}
	for k, v := range node.attributes {
		w.Write([]byte(" "))
		w.Write([]byte(k))
		w.Write([]byte(`="`))
		w.Write([]byte(v))
		w.Write([]byte(`"`))
	}
	w.Write([]byte(">"))
	cindent := indent + indentStep
	for _, c := range node.children {
		if c.text == "" {
			w.Write([]byte("\n"))
		}
		c.WriteTo(w, cindent, indentStep)
	}
	if len(node.children) > 0 {
		w.Write([]byte("\n"))
		w.Write([]byte(indent))
	}
	w.Write([]byte("</"))
	w.Write([]byte(node.tag))
	w.Write([]byte(">"))
	return nil
}

func (node *Node) ToString() (string, error) {
	var buffer bytes.Buffer
	err := node.WriteTo(&buffer, "", "  ")
	return buffer.String(), err
}

func parseTag(tag string) (string, string, []string, error) {
	const (
		sTag   = iota
		sID    = iota
		sClass = iota
	)
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
