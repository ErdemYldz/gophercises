package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link in a HTML
type Link struct {
	Href string
	Text string
}

var r io.Reader
var ll []Link

// Parse takes in an HTML and will return
// a slice of links from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	// var f func(*html.Node)
	// f = func(n *html.Node) {
	// 	if n.Type == html.ElementNode && n.Data == "a" {
	// 		l := Link{}
	// 		for _, a := range n.Attr {
	// 			if a.Key == "href" {
	// 				fmt.Println(a.Val)
	// 				l.Href = a.Val
	// 				break
	// 			}
	// 		}
	// 		// fmt.Printf("%#v\n", n)
	// 		// fmt.Printf("%#v\n", n.FirstChild)
	// 		fmt.Printf("%#v\n", n.FirstChild.Data)
	// 		l.Text = strings.TrimSpace(n.FirstChild.Data)
	// 		// fmt.Printf("%#v\n", n.FirstChild.FirstChild.Data)
	// 		// fmt.Printf("%#v\n", n.FirstChild.NextSibling.Data)
	// 		// // fmt.Printf("%#v\n", n.FirstChild.NextSibling.NextSibling)
	// 		// fmt.Printf("%#v\n", n.FirstChild.NextSibling.NextSibling.FirstChild.Data)
	// 		// // fmt.Printf("%#v\n", n.LastChild)
	// 		// // fmt.Printf("%#v\n", n.FirstChild.NextSibling)
	// 		// // fmt.Printf("%#v\n", n.FirstChild.NextSibling.FirstChild.Data)
	// 		// // fmt.Printf("%#v\n", n.FirstChild.NextSibling.NextSibling.Data)
	// 		// // fmt.Printf("%#v\n", n.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data)

	// 		ll = append(ll, l)
	// 	}
	// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 		f(c)
	// 	}
	// }
	// f(doc)
	// return ll, nil
	// dfs(doc, "")
	// fmt.Println(doc)
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		// fmt.Println(node)
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	if n.Attr != nil {
		for _, attr := range n.Attr {
			fmt.Println(attr.Key, attr.Val)
			ret.Href = attr.Val
		}
	}
	// if n.FirstChild != nil {
	// 	fmt.Println(n.FirstChild.Data)
	// 	ret.Text = n.FirstChild.Data
	// }
	str := text(n)
	ret.Text = str
	return ret
}

func text(n *html.Node) string {

	if n.Type == html.TextNode {
		return n.Data
	}

	// if n.Type != html.ElementNode {
	// 	return ""
	// }

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += " " + text(c)
	}
	fmt.Printf("%+v\n", strings.Fields(ret))
	ret = strings.Join(strings.Fields(ret), " ")
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
