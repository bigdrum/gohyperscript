package main

import (
	"github.com/bigdrum/gohyperscript/htmltoh"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("htmltoh", map[string]interface{}{
		"convert": htmltoh.HTMLToH,
	})
}
