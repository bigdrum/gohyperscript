package h

func MetaNameContent(name, content string) N {
	return H("meta", Attr("name", name), Attr("content", content))
}
