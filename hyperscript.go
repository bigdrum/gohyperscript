package h

import (
	"bytes"
	"fmt"
	"html"
	"strings"
)

type StringWriter interface {
	WriteString(s string) (int, error)
}

type Token interface {
	ICanBeToken()
}

type N interface {
	Token
	ToHTML(w StringWriter) error
}

type tagNode struct {
	tag  string
	toks []interface{}
}

type List []N

func (ns List) ICanBeToken() {}

func (ns List) ToHTML(w StringWriter) error {
	for _, n := range ns {
		err := n.ToHTML(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ns *List) Add(n N) {
	*ns = append(*ns, n)
}

func L(nlist ...N) List {
	return List(nlist)
}

type Attrs []KV

func (a Attrs) ICanBeToken() {}

type KV struct {
	K string
	V string
}

func Attr(k, v string) KV {
	return KV{K: k, V: v}
}

func (a KV) ICanBeToken() {}

type S string

func (s S) ICanBeToken() {}

func (s S) ToHTML(w StringWriter) error {
	w.WriteString(html.EscapeString(string(s)))
	return nil
}

func (s S) String() string {
	return string(s)
}

type DangerousUnescaped string

func (s DangerousUnescaped) ICanBeToken() {}

func (s DangerousUnescaped) ToHTML(w StringWriter) error {
	w.WriteString(string(s))
	return nil
}

func H(tag string, nodes ...interface{}) N {
	return tagNode{tag, nodes}
}

func (t tagNode) ICanBeToken() {}

func (t tagNode) ToHTML(w StringWriter) error {
	tag, id, classNames, err := parseTag(t.tag)
	if err != nil {
		return err
	}

	w.WriteString("<")
	w.WriteString(tag)

	if id != "" {
		w.WriteString(` id="`)
		w.WriteString(id)
		w.WriteString(`"`)
	}

	childrenOutputBegan := false
	for _, n := range t.toks {
		if n == nil {
			continue
		}
		switch cnode := n.(type) {
		case Attrs:
			if childrenOutputBegan {
				return fmt.Errorf("all attributes must be specified before children, violating attr: %v", cnode)
			}
			for _, attr := range cnode {
				if attr.K == "class" {
					if strings.ContainsAny(attr.V, `<>"`) {
						return fmt.Errorf("class name contains invalid character: %s", classNames)
					}
					classNames = append(classNames, attr.V)
				} else {
					w.WriteString(` `)
					w.WriteString(attr.K)
					w.WriteString(`="`)
					w.WriteString(html.EscapeString(attr.V))
					w.WriteString(`"`)
				}
			}
		case KV:
			if childrenOutputBegan {
				return fmt.Errorf("all attributes must be specified before children, violating attr: %v", cnode)
			}
			if cnode.K == "class" {
				if strings.ContainsAny(cnode.V, `<>"`) {
					return fmt.Errorf("class name contains invalid character: %s", classNames)
				}
				classNames = append(classNames, cnode.V)
			} else {
				w.WriteString(` `)
				w.WriteString(cnode.K)
				w.WriteString(`="`)
				w.WriteString(html.EscapeString(cnode.V))
				w.WriteString(`"`)
			}
		default:
			if childrenOutputBegan == false {
				if len(classNames) > 0 {
					w.WriteString(` class="`)
					w.WriteString(classNames[0])
					for _, c := range classNames[1:] {
						w.WriteString(" ")
						w.WriteString(c)
					}
					w.WriteString(`"`)
				}
				w.WriteString(">")
			}
			childrenOutputBegan = true

			switch cnode := n.(type) {
			case string:
				w.WriteString(html.EscapeString(cnode))
			case N:
				err := cnode.ToHTML(w)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("Invalid node %v", n)
			}
		}
	}
	if childrenOutputBegan {
		w.WriteString("</")
		w.WriteString(tag)
		w.WriteString(">")
	} else {
		if len(classNames) > 0 {
			w.WriteString(` class="`)
			w.WriteString(classNames[0])
			for _, c := range classNames[1:] {
				w.WriteString(" ")
				w.WriteString(c)
			}
			if isVoidElement(tag) {
				w.WriteString(`"/>`)
				return nil
			}
			w.WriteString(`"></`)
		} else {
			if isVoidElement(tag) {
				w.WriteString(`/>`)
				return nil
			}
			w.WriteString(`></`)
		}
		w.WriteString(tag)
		w.WriteString(">")
	}
	return nil
}

func ToString(n N) (string, error) {
	buf := bytes.Buffer{}
	err := n.ToHTML(&buf)
	if err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

func isVoidElement(s string) bool {
	// https://www.w3.org/TR/html5/syntax.html#void-elements
	switch s {
	case "area":
		return true
	case "base":
		return true
	case "br":
		return true
	case "col":
		return true
	case "embed":
		return true
	case "hr":
		return true
	case "img":
		return true
	case "input":
		return true
	case "keygen":
		return true
	case "link":
		return true
	case "meta":
		return true
	case "param":
		return true
	case "source":
		return true
	case "track":
		return true
	case "wbr":
		return true
	}
	return false
}
