package helpers

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func GetListElementByTag(n *html.Node, tag string) []*html.Node {
	if n == nil {
		return nil
	}
	var result []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == tag {
			result = append(result, c)
		}
	}
	return result
}

func GetAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func RenderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)

	err := html.Render(w, n)

	if err != nil {
		return ""
	}
	return buf.String()
}

// nolint:unused // This function used next turn
func checkId(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttribute(n, "id")
		if ok && s == id {
			return true
		}
	}
	return false
}

func checkClass(n *html.Node, class string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttribute(n, "class")
		if ok && strings.Contains(s, class) {
			return true
		}
	}
	return false
}

func traverse(n *html.Node, id string, fn func(node *html.Node, id string) bool) *html.Node {
	if fn(n, id) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res := traverse(c, id, fn)
		if res != nil {
			return res
		}
	}
	return nil
}

// nolint:deadcode,unused // This function used next turn
func GetElementById(n *html.Node, id string) *html.Node {
	return traverse(n, id, checkId)
}

func GetElementByClass(n *html.Node, class string) *html.Node {
	return traverse(n, class, checkClass)
}
