package main

import (
	"context"
	"fmt"
	"log"
	"nowcoder/logic"
	"os"
)

var logger *log.Logger

func main() {
	ctx := context.TODO()

	// 日志打印参考: https://studygolang.com/articles/1168
	logfile, err := os.OpenFile("/Users/maojianing/workspace/tmp/spider_log_20210729", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open log failed. err:%v \n", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	kewords := []string{"grpc"}
	for _, keyword := range kewords {
		for page := int64(1); page <= 5000; page++ {
			res, err := logic.HandleSearchPageLogic.GetSubUrl(ctx, keyword, page)
			if err != nil {
				log.Printf("keyword:%v page:%v err:%v\n", keyword, page, err)
				panic(err.Error())
			}
			log.Printf("res:%v\n", res)
			fmt.Printf("res:%v\n", res)
			return
		}
	}
}
