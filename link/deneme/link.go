package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

// type nodeType []Tag

// type Tag struct {
// 	tName  string
// 	tAttrs []Attributes
// 	tText  string
// }
// type Attributes map[string]string
type soup struct {
	n *html.Node
}

type node struct {
	*html.Node
}

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	// s := `<a href="/dog"><span>Something in a span</span>Text not in a span <b>Bold text!</b></a>`
	// 	s := `<html>
	// <body>
	//   <h1>Hello!</h1>
	//   <a href="/other-page">A link to another page</a>
	// </body>
	// </html>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("|%-12s|%-22s|%-22v|%-22v|%-22v|%-22v|%-22v\n", "where", "tag", "parent", "firstChild", "lastChild", "PreSibling", "NextSibling")

	// nodes := linkNodes(doc, "a")
	fmt.Printf("|%-12s|%-22s|%-22v|%-22v|%-22v|%-22v|%-22v\n", "where", "tag", "parent", "firstChild", "lastChild", "PreSibling", "NextSibling")
	// fmt.Println("len(nodes): ", len(nodes))
	// for i, node := range nodes {
	// 	// fmt.Printf("%+v\n", node)
	// 	printNode(node, fmt.Sprint(i))
	// }
	// fmt.Println(nodes[0])
	find := soup{n: doc}
	nnn := find.findAll("a")
	// printNode(nnn[0], "last method")
	for i, node := range nnn {
		// fmt.Printf("%+v\n", node)
		printNode(node, fmt.Sprint(i))
		
	}
}
func text(n *html.Node) string {
	var b bytes.Buffer
	html.Render(&b, n)
	return b.String()
}

func (s *soup) findAll(tag string) []*html.Node {

	return linkNodes(s.n, tag)
}

func linkNodes(n *html.Node, tag string) []*html.Node {
	printNode(n, "begin:")
	// fmt.Printf("%+v\n", n)
	if n.Type == html.ElementNode && n.Data == tag && n.Attr != nil {
		return []*html.Node{n}
	}

	var ret []*html.Node
	// if n.DataAtom != 0 && n.Data != "html" && n.Data != "head" {
	// 	ret = []*html.Node{n}
	// }
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printNode(c, "inside For:")
		// linkNodes(c)
		ret = append(ret, linkNodes(c, tag)...)
		// fmt.Println(ret)
	}

	printNode(n, "after For:")
	fmt.Println("************************************************************************************************************")
	return ret
}

func printNode(n *html.Node, custom string) {
	var tag, parent, fc, lc, ps, ns string

	tag = n.Data

	parent = "nil"
	if n.Parent != nil {
		parent = n.Parent.Data
	}

	fc = "nil"
	if n.FirstChild != nil {
		fc = n.FirstChild.Data
	}

	lc = "nil"
	if n.LastChild != nil {
		lc = n.LastChild.Data
	}

	ps = "nil"
	if n.PrevSibling != nil {
		ps = n.PrevSibling.Data
	}

	ns = "nil"
	if n.NextSibling != nil {
		ns = n.NextSibling.Data
	}

	fmt.Printf("|%-12s|%-22s|%-22v|%-22v|%-22v|%-22v|%-22v\n", custom, tag, parent, fc, lc, ps, ns)
}
