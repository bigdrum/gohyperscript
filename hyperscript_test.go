package hyperscript_test

import (
	"fmt"
	"testing"
)

import diff "github.com/kylelemons/godebug/diff"
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
		H(".entity", Style{"margin-top": "10px", "margin-bottom": "10px"}, "hoho"),
		func() *Node {
			n := H()
			for i := 0; i < 2; i++ {
				n.Add(H("p", fmt.Sprintf("p %d", i)))
			}
			return n
		}(),
		H("form.control",
			Style{"margin-top": "10px"},
			Attr{"data-haha": "kk", "class": "hello"}),
	)
	node = page("hello", node)
	s, err := node.ToString()
	if err != nil {
		t.Error(err)
	}
	expectedS := `<html lang="en">
  <head>
    <meta charset="utf-8"></meta>
    <title>hello</title>
  </head>
  <body>
    <div class="content" id="header">
      <h1>hello world</h1>
      <div class="entity">hoho</div>
      <p>p 0</p>
      <p>p 1</p>
      <form class="control hello" data-haha="kk"></form>
    </div>
  </body>
</html>`
	if s != expectedS {
		t.Error(s, diff.Diff(s, expectedS))
	}
}

func TestEscape(t *testing.T) {
	// This library doesn't support CSS escape.
	// We just test it won't go crazy.
	s, err := H("#test'.d'", Attr{"data-a": `"'<<>><script>`}).ToString()
	if err != nil {
		t.Error(err)
	}
	expectedS := `<div class="d&#39;" data-a="&#34;&#39;&lt;&lt;&gt;&gt;&lt;script&gt;" id="test&#39;"></div>`
	if s != expectedS {
		t.Error(diff.Diff(s, expectedS))
	}

	_, err = H("<#").ToString()
	if err == nil {
		t.Error("Expected error.")
	}

	_, err = H(">#").ToString()
	if err == nil {
		t.Error("Expected error.")
	}
}
