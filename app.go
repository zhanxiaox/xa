package xa

import (
	"fmt"
	"os"
	"slices"
)

type App struct {
	info    AppInfo
	cmds    map[string]*command
	runtime runtime
}

func New(info AppInfo) *App {
	return &App{
		info:    info,
		cmds:    make(map[string]*command),
		runtime: runtime{},
	}
}

type AppInfo struct {
	Name              string
	Desc              string
	Author            string
	Contact           string
	Version           string
	Usage             string
	EnableDefaultHelp bool
}

type runtime struct {
	Path string
	Cmd  string
	Args []string
}

func (a *App) GetRuntime() runtime {
	return a.runtime
}

func (a *App) HasArg(arg string) bool {
	return slices.Contains(a.runtime.Args, arg)
}

func (a *App) GetArg(arg string) string {
	index := slices.Index(a.runtime.Args, arg)
	if index > 0 && len(os.Args) >= index+1 {
		return a.runtime.Args[index+1]
	}
	return ""
}

func (a *App) GetAppInfo() AppInfo {
	return a.info
}

func (a *App) Run() {
	a.runtime.Path = os.Args[0]
	if len(os.Args) > 1 {
		a.runtime.Cmd = os.Args[1]
	}
	if len(os.Args) > 2 {
		a.runtime.Args = os.Args[2:]
	}
	if a.info.EnableDefaultHelp {
		a.Cmd("help", defaultHelp).Desc("Print this default help information")
	}
	if cmd, ok := a.cmds[a.runtime.Cmd]; ok {
		cmd.call(a)
	} else {
		if a.info.EnableDefaultHelp {
			defaultHelp(a)
		} else {
			fmt.Println("Unknown command:", a.runtime.Cmd)
		}
	}
}

func (a *App) Cmd(name string, call func(*App)) (c *command) {
	c = &command{Name: name, call: call}
	a.cmds[name] = c
	return
}

func (a *App) GetCmd(name string) command {
	return *a.cmds[name]
}

func defaultHelp(a *App) {
	fmt.Println(a.info.Name, a.info.Version)
	fmt.Println(a.info.Desc)
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println(a.info.Usage)
	fmt.Println()
	fmt.Println("OPTIONS:")
	for _, cmd := range a.cmds {
		fmt.Printf("%-17s %v", cmd.Name, cmd.Description)
		for _, argment := range cmd.Argments {
			fmt.Printf("\n  %-15s %v", argment.Name, argment.Description)
		}
		fmt.Println()
	}
}
