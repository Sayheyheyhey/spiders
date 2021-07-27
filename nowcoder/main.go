package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[rel=prefetch]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		fmt.Printf("url:%v\n", e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.nowcoder.com/search?type=post&order=time&query=grpc&subType=2&tagId=&page=1")
}
