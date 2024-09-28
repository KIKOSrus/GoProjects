package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"link"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type loc struct {
	Value string `xml:"loc"`
}

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func sortDomain(r *bytes.Reader, base string, links *map[string]int) []string {
	var recuLink []string
	doc, _ := html.Parse(r)
	hrefs := link.ParseLinks(doc)
	for _, l := range hrefs {
		switch {
		case strings.HasPrefix(l.Href, "/") && !strings.Contains(l.Href, "#"):
			var link string = base + l.Href
			if _, ok := (*links)[link]; !ok {
				(*links)[link] = 1
				recuLink = append(recuLink, link)
			}
		case strings.HasPrefix(l.Href, base) && !strings.Contains(l.Href, "#"):
			var link string = l.Href
			if _, ok := (*links)[link]; !ok {
				(*links)[link] = 1
				recuLink = append(recuLink, link)
			}
		}
	}
	return recuLink
}

func get(url string) (*bytes.Reader, string) {
	resp, err := http.Get(url)
	reqUrl := resp.Request.URL
	base := reqUrl.Scheme + "://" + reqUrl.Host
	link.Check(err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return bytes.NewReader(body), base
}

// говнокод да простит меня господь
// лень переделывать

func getRecursively(recuLink []string, limit int, links *map[string]int) {
	if limit > 0 {
		for _, l := range recuLink {
			reader, base := get(l)
			getRecursively(sortDomain(reader, base, links), limit-1, links)
		}
	}
}

func main() {
	limit := 3
	links := make(map[string]int)
	recuLinks := []string{"https://gophercises.com/"}

	getRecursively(recuLinks, limit, &links)
	var toXml urlset
	toXml.Xmlns = xmlns
	for l := range links {
		toXml.Urls = append(toXml.Urls, loc{Value: l})
	}
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", " ")
	fmt.Print(xml.Header)
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}

}
