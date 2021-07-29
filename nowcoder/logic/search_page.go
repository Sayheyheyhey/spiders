package logic

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"sync"
)

type SearchPageLogic struct {
}

var HandleSearchPageLogic = &SearchPageLogic{}

func (m *SearchPageLogic) GetSubUrl(ctx context.Context, keyword string, pageNum int64) ([]string, error) {

	res := make([]string, 0)
	var mu sync.Mutex
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[rel=prefetch]", func(e *colly.HTMLElement) {
		mu.Lock()
		url := e.Attr("href")
		res = append(res, url)
		mu.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	url := fmt.Sprintf("https://www.nowcoder.com/search?type=post&order=time&query=%s&subType=2&tagId=&page=%d", keyword, pageNum)
	c.Visit(url)
	log.Printf("get %d sub url in url:%v", len(res), url)

	return res, nil
}
