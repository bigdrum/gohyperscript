package h

func ClassName(className string, enabled bool) Token {
	if !enabled {
		return nil
	}
	return Attr("class", className)
}

func Javascript(js string) N {
	return H("script", Attr("type", "text/javascript"), DangerousUnescaped(js))
}

func ScopedStyle(style string) N {
	return H("style", Attr("scoped", "scoped"), S(style))
}

func Href(href string) Token {
	return Attr("href", href)
}

func TargetBlank() Token {
	return Attr("target", "_blank")
}

func Style(style string) Token {
	return Attr("style", style)
}

func Src(src string) Token {
	return Attr("src", src)
}
