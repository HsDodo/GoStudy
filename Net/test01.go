package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

//func main() {
//	start := time.Now()
//	ch := make(chan string) //创建一个通道
//	for _, url := range os.Args[1:] {
//		go fetch(url, ch) //并发执行 用go 来开启一个goroutine 也就是并发执行 ，注意这里没加go的话，就是顺序执行了，会很慢
//	}
//	for range os.Args[1:] {
//		fmt.Println(<-ch) //从通道ch中接收数据,取出数据 , 注意这里若是没有全部取出的话，会导致其他的goroutine还没结束 main函数就提前结束
//	}
//	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
//}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url) //resp 为响应数据流 类型为 *http.Response
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body) //io.Copy(dst, src) 从src中读取数据,写入到dst中,返回写入的字节数和遇到的错误, io.Discard是一个实现了Writer接口的变量,对它的任何写入都将被丢弃
	resp.Body.Close()                             //关闭resp.Body数据流,防止资源泄露
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()                  //计算从开始到结束的时间
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url) //将结果发送到通道ch中
}

func test01() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //resp : 数据响应流 类型为 *http.Response
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)

	}
}
