package xa

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type App struct {
	meta     Meta
	commands map[string]*Meta
}

type Meta struct {
	Name        string
	Description string
	Author      string
	Contact     string
	Version     string
	Usage       string
	Call        func(*App)
	Params      []Meta
}

func New(meta Meta) *App {
	return &App{
		meta:     meta,
		commands: make(map[string]*Meta),
	}
}

func (app *App) Run() {
	if len(os.Args) < 2 {
		fmt.Println("No command specified")
		return
	}
	if command, ok := app.commands[os.Args[1]]; ok {
		command.Call(app)
	} else {
		fmt.Println("Unknown command:", os.Args[1])
	}
}

func (app *App) Command(meta Meta) {
	app.commands[meta.Name] = &meta
}

func getArgs() []string {
	if len(os.Args) < 2 {
		return []string{}
	}
	return os.Args[2:]
}

func (app *App) GetArgsByIndex(index int) string {
	args := getArgs()
	if len(args) > index {
		return args[index]
	}
	return ""
}

func (app *App) GetArgsByName(name string) string {
	args := getArgs()
	if index := slices.Index(args, name); index >= 0 {
		if len(args)-1 > index+1 {
			return args[index+1]
		}
	}
	return ""
}

func (app *App) GetArgsByEqual(name string) string {
	args := getArgs()
	for _, arg := range args {
		_arg := strings.Split(arg, "=")
		if len(_arg) == 2 && _arg[0] == name {
			return _arg[1]
		}
	}
	return ""
}

func Help(app *App) {
	fmt.Println(app.meta.Name, app.meta.Version)
	fmt.Println(app.meta.Description)
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println(app.meta.Usage)
	fmt.Println()
	fmt.Println("OPTIONS:")
	for _, cmd := range app.commands {
		fmt.Printf("%-17s %v", cmd.Name, cmd.Description)
		for _, param := range cmd.Params {
			fmt.Printf("\n  %-15s %v", param.Name, param.Description)
		}
		fmt.Println()
	}
}
