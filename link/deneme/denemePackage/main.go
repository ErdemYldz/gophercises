package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ErdemYldz/gophercises/link/deneme/gosoup"

	"golang.org/x/net/html"
)

func main() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	// s := `<a href="/dog">
	// <span>Something in a span</span>
	// Text not in a span
	// <b>Bold text!</b>
	// </a>
	// `
	// s := `<html>
	// <body>
	//   <h1>Hello!</h1>
	//   <a href="/other-page">A link to another page</a>
	// </body>
	// </html>`

	// resp, err := http.Get("http://example.com/")
	// if err != nil {
	// 	log.Fatalln("error while requesting: ", err)
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln("error while reading the body: ", err)
	// }
	// fmt.Println("response body: ", string(body))

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}

	soup := gosoup.NewSoup(doc)

	fmt.Println(soup.RawText())

	nodes := soup.FindAll("body")
	for _, node := range nodes {
		fmt.Println(node.RawText())
		fmt.Println(node.Text())
		fmt.Println(node.FindAll("href"))
		fmt.Println(node.ParentSoup.RawText())
		fmt.Println(node.ParentSoup.FindAll("a")[0].RawText())
	}

	// fmt.Println(nodes[0].FindAll("span")[0].RawText())

}
