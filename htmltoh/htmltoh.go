package htmltoh

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func quoteText(buf *bytes.Buffer, text string) {
	if strconv.CanBackquote(text) {
		buf.WriteString("`")
		buf.WriteString(text)
		buf.WriteString("`")
		return
	}
	buf.WriteString(strconv.Quote(text))
}

func HTMLToH(source string) (string, error) {
	doc, err := html.Parse(strings.NewReader(source))
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer

	var outputNode func(n *html.Node, inTag bool)
	outputNode = func(n *html.Node, inTag bool) {
		walkChildren := func(inTag2 bool) {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				outputNode(c, inTag2)
			}
		}
		writeTagNode := func() {
			id := ""
			classNames := ""
			multiLine := func() bool {
				if n.FirstChild == nil {
					return false
				}
				if n.FirstChild.NextSibling != nil {
					return true
				}
				if n.FirstChild.Type == html.TextNode {
					return false
				}
				return true
			}()
			var attrs []html.Attribute
			for _, attr := range n.Attr {
				if attr.Key == "id" {
					id = attr.Val
					continue
				}
				if attr.Key == "class" {
					classNames = attr.Val
					continue
				}
				attrs = append(attrs, attr)
			}
			buf.WriteString(`h.H("`)
			buf.WriteString(n.Data)
			if id != "" {
				buf.WriteString("#")
				buf.WriteString(id)
			}
			if classNames != "" {
				for _, cls := range strings.Split(classNames, " ") {
					buf.WriteString(".")
					buf.WriteString(cls)
				}
			}
			buf.WriteString("\",")
			if multiLine {
				buf.WriteString("\n")
			}
			if len(attrs) == 1 {
				buf.WriteString("h.Attr(")
				buf.WriteString(strconv.Quote(attrs[0].Key))
				buf.WriteString(`, `)
				buf.WriteString(strconv.Quote(attrs[0].Val))
				buf.WriteString("),")
			} else if len(attrs) > 1 {
				buf.WriteString("h.Attrs{")
				for i, attr := range attrs {
					if i > 0 {
						buf.WriteString(",")
					}
					buf.WriteString(`{`)
					buf.WriteString(strconv.Quote(attr.Key))
					buf.WriteString(`, `)
					buf.WriteString(strconv.Quote(attr.Val))
					buf.WriteString(`}`)
				}
				buf.WriteString("},")
				if multiLine {
					buf.WriteString("\n")
				}
			}
			walkChildren(true)
			buf.WriteString(")")
			if inTag {
				buf.WriteString(",\n")
			}
		}

		switch n.Type {
		case html.DocumentNode:
			walkChildren(false)
			return
		case html.DoctypeNode:
			return
		case html.CommentNode:
			buf.WriteString("\n// ")
			buf.WriteString(n.Data)
			buf.WriteString("\n")
			return
		case html.ElementNode:
			writeTagNode()
			return
		case html.TextNode:
			text := strings.TrimSpace(n.Data)
			if text == "" {
				return
			}
			buf.WriteString("`")
			// if n.Data[0] == ' ' {
			// 	buf.WriteString(" ")
			// }
			buf.WriteString(text)
			// if n.Data[len(n.Data)-1] == ' ' {
			// 	buf.WriteString(" ")
			// }
			buf.WriteString("`")
			if inTag {
				buf.WriteString(",\n")
			}
			return
		default:
			panic(fmt.Errorf("Unknown node: %+v", n))
		}
	}
	outputNode(doc, false)
	dst := buf.Bytes()
	formatted, err := format.Source(dst)
	if err != nil {
		fmt.Println(string(dst))
		return "", err
	}
	return string(formatted), nil
}
