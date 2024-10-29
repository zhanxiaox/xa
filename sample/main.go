package main

import (
	"fmt"

	"github.com/zhanxiaox/xa"
)

func main() {
	a := xa.New()
	a.Load("download", "download file").Args("-p", "并发数").Args("-h", "帮助").Call(download)
	a.Load("upload", "upload file").Call(download)
}

func download(c *xa.Command) {
	fmt.Println(c.GetName(), c.GetDesc(), c.GetArgs(), c.Get("-p"))
}

func abb(a ...string) {

}
