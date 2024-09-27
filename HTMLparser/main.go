package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

// TODO: Создать struct Link {Href string, Text string}
// хранить в нем то что пропарсили
// возвращать массив Links
// научится игнорить комменты и победа!

type Link struct {
	Href string
	Text string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func parseText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		return parseText(c)
	}
	return ""
}

func BeautifulPrint(l Link) {
	fmt.Printf("Link{\n\tHref: \"%s\",\n\tText: \"%s\"\n}\n", l.Href, l.Text)
}

func parseLinks(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link := Link{attr.Val, parseText(n)}
				BeautifulPrint(link)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseLinks(c)
	}
}

func main() {
	filename := flag.String("html", "ex1.html", "the filename with html code")
	flag.Parse()
	file, err := os.Open("htmls/" + *filename)
	check(err)
	doc, _ := html.Parse(file)
	parseLinks(doc)

}
