package main

import (
	"fmt"

	"github.com/zhanxiaox/xa"
)

func main() {
	a := xa.New().Author("zx").Version("1.0.0").Desc("xa is a command line tool").Name("xa sample app")
	c := a.Load("a", "download file").Call(download).Desc("test").Args("-p", "test")
	fmt.Println("99999-", c.GetName(), c.GetDesc(), c.GetArgs(), c.Get("-p"))

	// a.Load("upload", "upload file").Call(download)
	// fmt.Printf("%+v", a)
	a.Run()
}

func download(c *xa.Command) {
	fmt.Println(c.GetName(), c.GetDesc(), c.GetArgs(), c.Get("-p"))
}
