package hyperscript_test

import "testing"

import . "github.com/bigdrum/gohyperscript"

func page(title string, content *Node) *Node {
	return H("html", Attr{"lang": "en"},
		H("head",
			H("meta", Attr{"charset": "utf-8"}),
			H("title", title)),
		H("body",
			content),
	)
}

func TestBasic(t *testing.T) {
	node := H("#header.content",
		H("h1", "hello world"),
		H(".entity", "hoho"),
		H("form.control",
			Attr{"data-haha": "kk", "class": "hello"},
			Style{}),
	)
	node = page("hello", node)
	s, err := node.ToString()
	if err != nil {
		t.Error(err)
	}
	if s != "" {
		t.Error(s)
	}
}
