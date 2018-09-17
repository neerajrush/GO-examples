package main

import (
	"fmt"
)

type WebPage struct {
	data  string
	links []string
}

type Crawler interface {
	Crawl(depth int) []string
}

type TestCrawler struct {
	VisitiedPagesMap map[string]WebPage
}

func (c *TestCrawler) Crawl(url string, depth int) {
	if len(url) == 0 || depth == 0 {
		return
	}
	if _,ok := c.VisitiedPagesMap[url]; !ok {
		if v, ok := TestDataMap[url]; ok {
			c.VisitiedPagesMap[url] = v
			for _,l := range v.links {
				c.Crawl(l, depth-1)
			}
		}
	}
}

func main() {
	T := TestCrawler{VisitiedPagesMap: make(map[string]WebPage),}
	T.Crawl("www.test1.com", 5)
	for k,v := range T.VisitiedPagesMap {
		fmt.Println(k, "==> WebPage## Data:", v.data, " Links:", v.links, "##" )
	}
}

var TestDataMap map[string]WebPage

func init() {
	TestDataMap = map[string]WebPage{ "www.test1.com" : WebPage{data: "This is test1 page", links: []string {"www.test1.com/link1", "www.test2.com", "www.test3.com"}},
				          "www.test2.com" : WebPage{data: "This is test2 page", links: []string {"www.test1.com", "www.test2.com/link2", "www.test3.com"}},
				          "www.test3.com" : WebPage{data: "This is test3 page", links: []string {"www.test1.com", "www.test2.com", "www.test3.com/link3"}}, }
}
