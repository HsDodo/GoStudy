package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	//构造一个waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	//设置需要等待处理
	//每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	//为每一个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		//启动一个goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	//启动一个goroutine 来监控是否所有的工作都做完了
	go func() {
		//等候所有任务完成
		waitGroup.Wait()

		//用关闭通道的方式，通知Display函数
		//可以退出程序了
		close(results)
	}()

	//启动函数，显示返回的结果，并且在最后一个结果显示完后返回
	Display(results)

}
