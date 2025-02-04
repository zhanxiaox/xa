package cmds

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"
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

func getUsableDomain() (string, error) {
	domains := []string{"https://golang.google.com/dl", "https://go.dev/dl", "https://golang.google.cn/dl"}
	for _, domain := range domains {
		go ping(domain)
	}
	select {
	case domain := <-complete:
		return domain, nil
	case <-time.After(5 * time.Second):
		return "", errors.New("timeout")
	}
}

func Update(app *xa.App) {
	if domain, err := getUsableDomain(); err != nil {
		fmt.Println(err.Error())
	} else {
		resp, err := http.Get(domain)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		html := string(body)
		re := regexp.MustCompile(`<a class="download downloadBox" href="([^"]+)">`)
		match := re.FindAllStringSubmatch(html, -1)
		hrefs := []string{}
		if len(match) > 0 {
			for _, m := range match {
				hrefs = append(hrefs, m[1])
			}
		} else {
			fmt.Println("not found")
			return
		}
		downloadUrl := ""
		for _, href := range hrefs {
			if strings.Contains(href, runtime.GOOS) {
				u, err := url.Parse(domain)
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				downloadUrl = u.Scheme + "://" + u.Hostname() + href
				break
			}
		}
		if downloadUrl == "" {
			fmt.Println("not found")
			return
		}
		fmt.Println(downloadUrl)
	}
}
