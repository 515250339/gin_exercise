package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// 获取当前时间
	start := time.Now()
	// make函数创建了一个传递string类型参数的channel
	ch := make(chan string)
	// for循环命令行参数
	for _, url := range os.Args[1:] {
		// 开启一个goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 接收并打印channel,for循环不需要key value
		fmt.Println(<-ch)
	}
	// main函数的执行时间
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// 参数类型:string , chan<- string
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// 把内容扔掉,只获取字节数
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	// 记录执行的秒数
	secs := time.Since(start).Seconds()
	// 发送给channel
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
