package htmltoh

import "testing"

func init() {
	// flag.Set("logtostderr", "true")
}

func TestBasic(t *testing.T) {
	cases := []struct {
		src string
		dst string
	}{
		{
			src: `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width">
    <title>ctrlq
    </title>
  </head>
  <body>
    <h1>Hello World
    </h1>
  </body>
</html> `,
			dst: `_none_empty`,
		},
		{
			src: `<!DOCTYPE html>
<html lang="en">
  <head profile="http://a9.com/-/spec/opensearch/1.1/">
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.1/css/bootstrap.min.css" rel="stylesheet">
    <link href="/-/site.css?v=83fa968c50840a0c43e964f10f0a754bc9fb77f1" rel="stylesheet">
    <title>format - GoDoc
    </title>
    <meta name="twitter:title" content="Package format">
    <meta property="og:title" content="Package format">
    <meta name="description" content="Package format implements standard formatting of Go source.">
    <meta name="twitter:description" content="Package format implements standard formatting of Go source.">
    <meta property="og:description" content="Package format implements standard formatting of Go source.">
    <meta name="twitter:card" content="summary">
    <meta name="twitter:site" content="@golang">
  </head>
  <body>
    <nav class="navbar navbar-default" role="navigation">
      <div class="container">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
            <span class="sr-only">Toggle navigation
            </span>
            <span class="icon-bar">
            </span>
            <span class="icon-bar">
            </span>
            <span class="icon-bar">
            </span>
          </button>
          <a class="navbar-brand" href="/">
            <strong>GoDoc
            </strong>
          </a>
        </div>
        <div class="collapse navbar-collapse">
          <ul class="nav navbar-nav">
            <li>
              <a href="/">Home
              </a>
            </li>
            <li>
              <a href="/-/index">Index
              </a>
            </li>
            <li>
              <a href="/-/about">About
              </a>
            </li>
          </ul>
          <form class="navbar-nav navbar-form navbar-right" id="x-search" action="/" role="search">
            <input class="form-control" id="x-search-query" type="text" name="q" placeholder="Search">
          </form>
        </div>
      </div>
    </nav>
    <div class="container">
      <div class="clearfix" id="x-projnav">
        <a href="/-/go">Go:
        </a>
        <a href="/go">go
        </a>
        <span class="text-muted">/
        </span>
        <span class="text-muted">format
        </span>
        <span class="pull-right">
          <a href="#pkg-index">Index
          </a>
          <span class="text-muted">|
          </span>
          <a href="#pkg-files">Files
          </a>
        </span>
      </div>
      <h2 id="pkg-overview">package format
      </h2>
      <p>
        <code>import "go/format"
        </code>
      <p>
        Package format implements standard formatting of Go source.
      </p>
      <h3 id="pkg-index" class="section-header">Index
        <a class="permalink" href="#pkg-index">&para;
        </a>
      </h3>
      <ul class="list-unstyled">
        <li>
          <a href="#Node">func Node(dst io.Writer, fset *token.FileSet, node interface{}) error
          </a>
        </li>
        <li>
          <a href="#Source">func Source(src []byte) ([]byte, error)
          </a>
        </li>
      </ul>
      <span id="pkg-examples">
      </span>
      <h4 id="pkg-files">
        <a href="https://golang.org/src/go/format/">Package Files
        </a>
        <a class="permalink" href="#pkg-files">&para;
        </a>
      </h4>
      <p>
        <a href="https://golang.org/src/go/format/format.go">format.go
        </a>
        <a href="https://golang.org/src/go/format/internal.go">internal.go
        </a>
      </p>
      <h3 id="Node" data-kind="f">func
        <a title="View Source" href="https://golang.org/src/go/format/format.go#L33">Node
        </a>
        <a class="permalink" href="#Node">&para;
        </a>
      </h3>
      <div class="funcdecl decl">
        <a title="View Source" href="https://golang.org/src/go/format/format.go#L33">❖
        </a>
        <pre>func Node(dst
<a href="/io">io
</a>.
<a href="/io#Writer">Writer
</a>, fset *
<a href="/go/token">token
</a>.
<a href="/go/token#FileSet">FileSet
</a>, node interface{})
<a href="/builtin#error">error
</a>
</pre>
      </div>
      <p>
        Node formats node in canonical gofmt style and writes the result to dst.
      </p>
      <p>
        The node type must be *ast.File, *printer.CommentedNode, []ast.Decl,
        []ast.Stmt, or assignment-compatible to ast.Expr, ast.Decl, ast.Spec,
        or ast.Stmt. Node does not modify node. Imports are not sorted for
        nodes representing partial source files (i.e., if the node is not an
        *ast.File or a *printer.CommentedNode not wrapping an *ast.File).
      </p>
      <p>
        The function may return early (before the entire result is written)
        and return a formatting error, for instance due to an incorrect AST.
      </p>
      <h3 id="Source" data-kind="f">func
        <a title="View Source" href="https://golang.org/src/go/format/format.go#L82">Source
        </a>
        <a class="permalink" href="#Source">&para;
        </a>
      </h3>
      <div class="funcdecl decl">
        <a title="View Source" href="https://golang.org/src/go/format/format.go#L82">❖
        </a>
        <pre>func Source(src []
<a href="/builtin#byte">byte
</a>) ([]
<a href="/builtin#byte">byte
</a>,
<a href="/builtin#error">error
</a>)
</pre>
      </div>
      <p>
        Source formats src in canonical gofmt style and returns the result
        or an (I/O or syntax) error. src is expected to be a syntactically
        correct Go source file, or a list of Go declarations or statements.
      </p>
      <p>
        If src is a partial source file, the leading and trailing space of src
        is applied to the result (such that it has the same leading and trailing
        space as src), and the result is indented by the same amount as the first
        line of src containing code. Imports are not sorted for partial source files.
      </p>
      <div id="x-pkginfo">
        <form name="x-refresh" method="POST" action="/-/refresh">
          <input type="hidden" name="path" value="go/format">
        </form>
        <p>Package format imports
          <a href="?imports">8 packages
          </a> (
          <a href="?import-graph">graph
          </a>) and is imported by
          <a href="?importers">442 packages
          </a>.
          Updated
          <span class="timeago" title="2016-04-21T19:22:31Z">2016-04-21
          </span>.
          <a href="javascript:document.getElementsByName('x-refresh')[0].submit();" title="Refresh this page from the source.">Refresh now
          </a>.
          <a href="?tools">Tools
          </a> for package owners.
      </div>
      <div id="x-jump" tabindex="-1" class="modal">
        <div class="modal-dialog">
          <div class="modal-content">
            <div class="modal-header">
              <h4 class="modal-title">Jump to identifier
              </h4>
              <br class="clearfix">
              <input id="x-jump-filter" class="form-control" autocomplete="off" type="text">
            </div>
            <div id="x-jump-body" class="modal-body" style="height: 260px; overflow: auto;">
              <div id="x-jump-list" class="list-group" style="margin-bottom: 0;">
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn" data-dismiss="modal">Close
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div id="x-footer" class="clearfix">
      <div class="container">
        <a href="https://github.com/golang/gddo/issues">Website Issues
        </a>
        <span class="text-muted">|
        </span>
        <a href="http://golang.org/">Go Language
        </a>
        <span class="pull-right">
          <a href="#">Back to top
          </a>
        </span>
      </div>
    </div>
    <div id="x-shortcuts" tabindex="-1" class="modal">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;
            </button>
            <h4 class="modal-title">Keyboard shortcuts
            </h4>
          </div>
          <div class="modal-body">
            <table>
              <tr>
                <td align="right">
                  <b>?
                  </b>
                </td>
                <td> : This menu
                </td>
              </tr>
              <tr>
                <td align="right">
                  <b>/
                  </b>
                </td>
                <td> : Search site
                </td>
              </tr>
              <tr>
                <td align="right">
                  <b>f
                  </b>
                </td>
                <td> : Jump to identifier
                </td>
              </tr>
              <tr>
                <td align="right">
                  <b>g
                  </b> then
                  <b>g
                  </b>
                </td>
                <td> : Go to top of page
                </td>
              </tr>
              <tr>
                <td align="right">
                  <b>g
                  </b> then
                  <b>b
                  </b>
                </td>
                <td> : Go to end of page
                </td>
              </tr>
              <tr>
                <td align="right">
                  <b>g
                  </b> then
                  <b>i
                  </b>
                </td>
                <td> : Go to index
                </td>
              </tr>
              <tr>
                <td align="right">
                  <b>g
                  </b> then
                  <b>e
                  </b>
                </td>
                <td> : Go to examples
                </td>
              </tr>
            </table>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn" data-dismiss="modal">Close
            </button>
          </div>
        </div>
      </div>
    </div>
    <script src="//code.jquery.com/jquery-2.0.3.min.js">
    </script>
    <script src="//maxcdn.bootstrapcdn.com/bootstrap/3.3.1/js/bootstrap.min.js">
    </script>
    <script src="/-/site.js?v=371de731c18d91c499d90b1ab0bf39ecf66d6cf7">
    </script>
    <script type="text/javascript">
      var _gaq = _gaq || [];
      _gaq.push(['_setAccount', 'UA-11222381-8']);
      _gaq.push(['_trackPageview']);
      (function() {
        var ga = document.createElement('script');
        ga.type = 'text/javascript';
        ga.async = true;
        ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') + '.google-analytics.com/ga.js';
        var s = document.getElementsByTagName('script')[0];
        s.parentNode.insertBefore(ga, s);
      }
      )();
    </script>
  </body>
</html>
`,
			dst: "_none_empty",
		},
		{
			src: `<div class=" hi  hello   kk  ">pp</div>`,
			dst: `h.H("div.hi.hello.kk", ` + "`pp`)",
		},
		{
			src: `<html>`,
		},
		{
			src: "<p>`h\"i`</p>",
			dst: "h.H(\"p\", \"`h\\\"i`\")",
		},
	}
	for i, c := range cases {
		hcode, err := HTMLToH(c.src)
		if err != nil {
			t.Error(i, err)
			continue
		}
		if c.dst == "_none_empty" {
			if hcode == "" {
				t.Error("cannot convert, ", i, c.src)
			}
		} else if hcode != c.dst {
			t.Error(i, c.dst, hcode)
		}
	}
}
