package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 循环命令行参数
	for _, url := range os.Args[1:] {
		//加入前缀 当参数不为 http:// 开头 则进行拼接
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		// 获取当前 url
		resp, err := http.Get(url)
		// 异常捕获
		if err != nil {
			//向标准错误流打印信息
			fmt.Fprintf(os.Stderr, "fetch: v%\n", err)
			//终止进程
			os.Exit(1)
		}
		// 获取 resp 中 Body 信息
		// b, err := ioutil.ReadAll(resp.Body)
		// 避免申请一个缓冲区,直接到标准输出流 从 resp.Body 读，并且写入 os.Stdout
		io.Copy(os.Stdout, resp.Body)
		// 关闭 Body 避免资源泄露
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("")
		// 输出 code 码
		fmt.Printf("%s", resp.Status)
	}
}
