package hyperscript

import (
	"bytes"
	"html"
	"io"
	"sort"
)

func sortedAttributes(m map[string]string) []string {
	l := make([]string, len(m))
	i := 0
	for k := range m {
		l[i] = k
		i++
	}
	sort.Strings(l)
	return l
}

func sortedClassNames(m map[string]bool) []string {
	l := make([]string, len(m))
	i := 0
	for k, v := range m {
		if !v {
			continue
		}
		l[i] = k
		i++
	}
	l = l[:i]
	sort.Strings(l)
	return l
}

// WriteIndent encodes the given dom tree into HTML.
func (node *Node) WriteIndent(w io.Writer, indent string, indentStep string) error {
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
		for _, c := range sortedClassNames(node.classNames) {
			w.Write(sep)
			w.Write([]byte(c))
			sep = byteSpace
		}
		w.Write([]byte(`"`))
	}

	attributeKeys := sortedAttributes(node.attributes)
	for _, k := range attributeKeys {
		w.Write([]byte(" "))
		w.Write([]byte(k))
		w.Write([]byte(`="`))
		w.Write([]byte(node.attributes[k]))
		w.Write([]byte(`"`))
	}
	w.Write([]byte(">"))
	cindent := indent + indentStep
	for _, c := range node.children {
		if c.text == "" {
			w.Write([]byte("\n"))
		}
		c.WriteIndent(w, cindent, indentStep)
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

// ToString converts the dom tree into HTML.
func (node *Node) ToString() (string, error) {
	var buffer bytes.Buffer
	err := node.WriteIndent(&buffer, "", "  ")
	return buffer.String(), err
}
