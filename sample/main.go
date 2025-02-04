package main

import (
	"github.com/zhanxiaox/xa"
	"github.com/zhanxiaox/xa/sample/cmds"
)

func main() {
	app := xa.New().Info(xa.AppInfo{
		Name:              "goup",
		Author:            "zhanxiaox",
		Version:           "0.0.1",
		Desc:              "goup is go toolchain installer",
		Usage:             "goup .exe [OPTIONS]",
		Contact:           "346084070@qq.com",
		EnableDefaultHelp: true,
	})
	app.NewCmd("update", cmds.Update).Desc("Update go latest version")
	app.NewCmd("version", cmds.Version).Desc("Print version information")
	app.NewCmd("install", cmds.Install).Desc("Install goup into Golang's system path (need root permisson)")
	app.NewCmd("uninstall", cmds.Uninstall).Desc("Remove goup from Golang's system path (need root permisson)")
	app.Run()
}
