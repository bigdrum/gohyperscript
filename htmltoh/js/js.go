package main

import (
	"bitbucket.org/applysquare/applysquare-go/pkg/util/h/htmltoh"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("htmltoh", map[string]interface{}{
		"convert": htmltoh.HTMLToH,
	})
}
