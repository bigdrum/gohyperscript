package h_test

import (
	"fmt"
	"strings"
	"testing"
)

import diff "github.com/kylelemons/godebug/diff"
import . "github.com/bigdrum/gohyperscript"

func page(title string, content N) N {
	return H("html", Attr("lang", "en"),
		H("head",
			H("meta", Attr("charset", "utf-8")),
			H("title", title)),
		H("body",
			content),
	)
}

func TestBasic(t *testing.T) {
	node := H("#header.content",
		H("h1", "hello world"),
		H(".entity", Attr("style", `margin-top: 10px; margin-bottom: 10px`), "hoho"),
		func() N {
			n := L()
			for i := 0; i < 2; i++ {
				n.Add(H("p", fmt.Sprintf("p %d", i)))
			}
			return n
		}(),
		H("form.control",
			Attrs{{"data-haha", "kk"}, {"class", "hello"}, {"style", `margin-top: 10px`}}),
	)
	node = page("hello", node)
	s, err := ToString(node)
	if err != nil {
		t.Error(err)
	}
	s = strings.Replace(s, ">", ">\n", -1)
	expectedS := `<html lang="en">
<head>
<meta charset="utf-8"/>
<title>
hello</title>
</head>
<body>
<div id="header" class="content">
<h1>
hello world</h1>
<div style="margin-top: 10px; margin-bottom: 10px" class="entity">
hoho</div>
<p>
p 0</p>
<p>
p 1</p>
<form data-haha="kk" style="margin-top: 10px" class="control hello">
</form>
</div>
</body>
</html>
`
	if s != expectedS {
		t.Error(s, diff.Diff(s, expectedS))
	}
}

func TestSingleSpan(t *testing.T) {
	s, err := ToString(H("span"))
	if err != nil {
		t.Error(err)
	}
	if s != "<span></span>" {
		t.Error(s)
	}
}

func TestVoidElements(t *testing.T) {
	s, err := ToString(H("img#ttt.kkk.aaa"))
	if err != nil {
		t.Error(err)
	}
	if s != `<img id="ttt" class="kkk aaa"/>` {
		t.Error(s)
	}
}

func TestEscape(t *testing.T) {
	// This library doesn't support CSS escape.
	// We just test it won't go crazy.
	s, err := ToString(H("#test'.d'", Attr("data-a", `"'<<>><script>`)))
	if err != nil {
		t.Error(err)
	}
	expectedS := `<div id="test'" data-a="&#34;&#39;&lt;&lt;&gt;&gt;&lt;script&gt;" class="d'"></div>`
	if s != expectedS {
		t.Error(diff.Diff(s, expectedS))
	}

	_, err = ToString(H("<#"))
	if err == nil {
		t.Error("Expected error.")
	}

	_, err = ToString(H(">#"))
	if err == nil {
		t.Error("Expected error.")
	}
	_, err = ToString(H("\"#"))
	if err == nil {
		t.Error("Expected error.")
	}
}
