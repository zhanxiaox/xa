package main

import (
	"gitee.com/zhanxiaox/xa"
	"gitee.com/zhanxiaox/xa/sample/cmds"
)

func main() {
	app := xa.New(xa.Meta{
		Name:        "goup",
		Author:      "zhanxiaox",
		Version:     "0.0.1",
		Description: "goup is go toolchain installer",
		Usage:       "goup .exe [OPTIONS]",
		Contact:     "346084070@qq.com",
	})
	app.Command("update", xa.Meta{
		Call:        cmds.Update,
		Description: "Update go latest version",
		Params: []xa.Meta{
			{Name: "-f", Description: "Force update"},
			{Name: "-v", Description: "Verbose output"},
			{Name: "-y", Description: "Skip confirmation"},
		},
	})
	app.Command("version", xa.Meta{Call: cmds.Version, Description: "Print version information"})
	app.Command("install", xa.Meta{Call: cmds.Install, Description: "Install goup into Golang's system path (need root permisson)"})
	app.Command("uninstall", xa.Meta{Call: cmds.Uninstall, Description: "Remove goup from Golang's system path (need root permisson)"})
	app.Command("help", xa.Meta{Call: xa.Help, Description: "xa default help"})
	app.Run()
}
