package main

import (
	"fmt"

	"github.com/zhanxiaox/xa"
)

func main() {
	a := xa.GetApp().Author("zhanxiaox").Version("1.0.0").Desc("xa is a command line tool").Name("xa sample app")
	a.Load("download", download).Desc("test").Args("-p", "test")
	a.Run()
}

func download(r *xa.Runtime) {
	fmt.Println(xa.GetApp().GetAuthor(), r.GetName(), r.HasArg("-p"))
}
