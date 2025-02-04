package xa

import (
	"fmt"
	"os"
	"slices"
)

type App struct {
	info    AppInfo
	cmds    map[string]*Command
	runtime runtime
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

func (a *App) SetRuntime() {
	a.runtime.Path = os.Args[0]
	if len(os.Args) > 1 {
		a.runtime.Cmd = os.Args[1]
	}
	if len(os.Args) > 2 {
		a.runtime.Args = os.Args[2:]
	}
}

func (a *App) GetRuntime() runtime {
	return a.runtime
}

func New() *App {
	return &App{cmds: make(map[string]*Command)}
}

func (a *App) Info(ai AppInfo) *App {
	a.info = ai
	return a
}

func (a *App) HasArg(arg string) bool {
	return slices.Contains(a.runtime.Args, arg)
}

func (a *App) GetArg(arg string) string {
	index := slices.Index(a.runtime.Args, arg)
	if index > 0 && len(os.Args) >= index+1 {
		return os.Args[index+1]
	}
	return ""
}

func (a *App) GetAppInfo() AppInfo {
	return a.info
}

func (a *App) Run() {
	a.SetRuntime()
	if a.info.EnableDefaultHelp {
		a.NewCmd("help", defaultHelp).Desc("Print this default help information")
	}
	if cmd, ok := a.cmds[a.runtime.Cmd]; ok {
		cmd.call(a)
	} else {
		if a.info.EnableDefaultHelp {
			defaultHelp(a)
		}
	}
}

func (a *App) NewCmd(cmd string, call func(*App)) *Command {
	c := Command{name: cmd, call: call}
	a.cmds[cmd] = &c
	return &c
}

func defaultHelp(a *App) {
	fmt.Println(a.info.Name, a.info.Version)
	fmt.Println(a.info.Desc)
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println(a.info.Usage)
	fmt.Println()
	fmt.Println("OPTIONS:")
	for _, v := range a.cmds {
		fmt.Printf("%-17s %v", v.name, v.desc)
		for _, v := range v.args {
			fmt.Printf("\n  %-15s %v", v.name, v.desc)
		}
		fmt.Println()
	}
}
