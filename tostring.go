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
func (node *Node) WriteIndent(w io.Writer, indent []byte, indentStep []byte) error {
	if node.err != nil {
		return node.err
	}
	if node.text != "" {
		w.Write([]byte(html.EscapeString(node.text)))
		return nil
	}

	cindent := indent
	if node.tag != "" {
		w.Write(indent)
		w.Write([]byte("<"))
		w.Write([]byte(node.tag))

		if len(node.classNames) > 0 {
			w.Write([]byte(` class="`))
			sep := []byte{}
			for _, c := range sortedClassNames(node.classNames) {
				w.Write(sep)
				w.Write([]byte(html.EscapeString(c)))
				sep = byteSpace
			}
			w.Write([]byte(`"`))
		}

		attributeKeys := sortedAttributes(node.attributes)
		for _, k := range attributeKeys {
			w.Write(byteSpace)
			w.Write([]byte(k))
			w.Write([]byte(`="`))
			w.Write([]byte(html.EscapeString(node.attributes[k])))
			w.Write([]byte(`"`))
		}
		w.Write([]byte(">"))
		cindent = append(indent, indentStep...)
	}

	for i, c := range node.children {
		if i == 0 && c.text == "" && node.tag != "" {
			w.Write([]byte("\n"))
		}
		c.WriteIndent(w, cindent, indentStep)
		if c.text == "" && (node.tag != "" || i != len(node.children)-1) {
			w.Write([]byte("\n"))
			if i == len(node.children)-1 {
				w.Write(indent)
			}
		}
	}

	if node.tag != "" {
		w.Write([]byte("</"))
		w.Write([]byte(node.tag))
		w.Write([]byte(">"))
	}
	return nil
}

// ToString converts the dom tree into HTML.
func (node *Node) ToString() (string, error) {
	var buffer bytes.Buffer
	err := node.WriteIndent(&buffer, []byte{}, []byte("  "))
	return buffer.String(), err
}
