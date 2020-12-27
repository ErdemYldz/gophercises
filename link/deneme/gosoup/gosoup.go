package gosoup

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

// Soup object holds the root node
type Soup struct {
	node       *html.Node
	ParentSoup *Soup
}

// NewSoup initializes and returns the Soup object
func NewSoup(doc *html.Node) *Soup {
	return &Soup{node: doc}
}

// func (s *Soup) getParent() *Soup {
// 	return &Soup{node: s.parentNode}

// }

// RawText returns the raw text of the node
func (s *Soup) RawText() string {
	var b bytes.Buffer
	html.Render(&b, s.node)
	return b.String()
}

// FindAll returns the nodes with spesified tag as a parameter
func (s *Soup) FindAll(tag string) []*Soup {
	return linkNodes(s.node, tag)
}

// Text returns the text from an HTML node
func (s *Soup) Text() string {

	if s.node.Type == html.TextNode {
		return s.node.Data
	}

	var ret string
	for c := s.node.FirstChild; c != nil; c = c.NextSibling {
		s := Soup{node: c}
		ret += " " + s.Text()
	}

	// fmt.Printf("%+v\n", strings.Fields(ret))
	ret = strings.Join(strings.Fields(ret), " ")
	return ret
}

func linkNodes(n *html.Node, tag string) []*Soup {
	if n.Type == html.ElementNode && n.Data == tag {
		s := &Soup{node: n, ParentSoup: &Soup{node: n.Parent}}
		return []*Soup{s}
	}

	var ret []*Soup
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c, tag)...)
	}
	return ret
}
