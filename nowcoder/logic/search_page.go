package logic

import (
	"context"
	"fmt"
	"github.com/gocolly/colly"
	"sync"
)

type SearchPageLogic struct {
}

var HandleSearchPageLogic = &SearchPageLogic{}

func (m *SearchPageLogic) GetSubUrl(ctx context.Context, keyword string, pageNum int64) ([]string, error) {
	c := colly.NewCollector()

	var mu sync.Mutex

	// Find and visit all links
	c.OnHTML("a[rel=prefetch]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.nowcoder.com/search?type=post&order=time&query=grpc&subType=2&tagId=&page=1")

}
