package cmds

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zhanxiaox/xa"
)

var complete = make(chan string, 1)

func ping(domain string) {
	resp, e := http.Get(domain)
	if e == nil && resp.StatusCode == 200 {
		complete <- domain
	}
}

func getUsableDomain() {
	domains := []string{"https://golang.google.com", "https://go.dev", "https://golang.google.cn"}
	for _, domain := range domains {
		go ping(domain)
	}
	select {
	case domain := <-complete:
		fmt.Println("可用域名：", domain)
	case <-time.After(5 * time.Second):
		fmt.Println("超时")
	}
}

func Update(app *xa.App) {
	fmt.Println(app.GetAuthor(), app.GetRuntime().Cmd)
	getUsableDomain()
}
