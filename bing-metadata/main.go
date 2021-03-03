package main

import (
	"archive/zip"
	"bytes"
	"clients/bing-metadata/metadata"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func handler(index int, search *goquery.Selection) {

	url, ok := search.Find("a").Attr("href")
	if !ok {
		return
	}
	fmt.Printf("%d: %s\n", index, url)

	result, err := http.Get(url)
	if err != nil {
		return
	}

	buf, err := ioutil.ReadAll(result.Body)

	if err != nil {
		return
	}

	defer result.Body.Close()

	r, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return
	}
	cp, ap, err := metadata.NewProperties(r)

	if err != nil {
		return
	}

	log.Printf(
		"%25s %25s - %s %s\n", cp.Creator, cp.LastModifiedBy, ap.Application, ap.GetMajorVersion())
}

func main() {
	domain := "nytimes.com"
	filetype := "docx"
	q := fmt.Sprintf(
		"site:%s && filetype:%s && instreamset:(url title):%s", domain,
		filetype,
		filetype)
	search := fmt.Sprintf("http://www.bing.com/search?q=%s", url.QueryEscape(q))
	doc, err := goquery.NewDocument(search)
	if err != nil {
		log.Panicln(err)
	}
	s := "html body div#b_content ol#b_results li.b_algo h2"
	doc.Find(s).Each(handler)
}
