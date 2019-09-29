package truncate

import (
	"strings"

	"golang.org/x/net/html"
)

// HTML truncates the input HTML snippet h to l runes.
func HTML(h string, l int) (string, error) {
	// wrapped by <html><head></head><body>...</body></html>
	d, err := html.Parse(strings.NewReader(h))
	if err != nil {
		return "", err
	}

	_ = node(d, l)

	// strip <html><head></head><body>...</body></html> and render
	var w strings.Builder
	for n := d.LastChild.LastChild.FirstChild; n != nil; n = n.NextSibling {
		if err := html.Render(&w, n); err != nil {
			return "", err
		}
	}

	return w.String(), nil
}

func node(n *html.Node, l int) int {
	if l <= 0 {
		n.Parent.RemoveChild(n)
		return l
	}

	if n.Type == html.TextNode {
		r := []rune(n.Data)
		if l < len(r) {
			r = r[:l]
			r[l-1] = '…'
			n.Data = string(r)
			return 0
		}

		return l - len(n.Data)
	}

	var cs []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		cs = append(cs, c)
	}

	for _, c := range cs {
		l = node(c, l)
	}

	return l
}
