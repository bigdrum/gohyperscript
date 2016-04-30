package h

import (
	"fmt"
	"strings"
)

func parseTag(tag string) (string, string, []string, error) {
	if strings.ContainsAny(tag, `<>"`) {
		return "", "", nil, fmt.Errorf("special charactor not supported for tag/class/id: %s", tag)
	}

	tagEnd := -1
	idStart := -1
	id := ""
	classStart := -1
	var classNames []string

	tlen := len(tag)
	for i, c := range tag {
		if c == '#' {
			if tagEnd == -1 {
				tagEnd = i
			}
			if idStart >= 0 {
				return "", "", nil, fmt.Errorf("id specified more than once: %s", tag)
			}
			if classStart >= 0 {
				classNames = append(classNames, tag[classStart:i])
				classStart = -1
			}
			idStart = i + 1
			continue
		}
		if c == '.' {
			if tagEnd == -1 {
				tagEnd = i
			}
			if idStart >= 0 && id == "" {
				id = tag[idStart:i]
			}
			if classStart >= 0 {
				classNames = append(classNames, tag[classStart:i])
			}
			classStart = i + 1
		}
	}
	if classStart >= 0 {
		classNames = append(classNames, tag[classStart:tlen])
	} else if idStart >= 0 && id == "" {
		id = tag[idStart:tlen]
	} else {
		tagEnd = tlen
	}

	pureTag := "div"
	if tagEnd > 0 {
		pureTag = tag[:tagEnd]
	}
	return pureTag, id, classNames, nil
}
