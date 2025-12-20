package xa

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type App struct {
	meta     Meta
	commands []command
}

type Meta struct {
	Name        string
	Description string
	Author      string
	Contact     string
	Version     string
	Usage       string
	Params      []Meta
}

type command struct {
	Meta
	Key  string
	Call func(App)
}

func New() *App {
	return &App{}
}

func (app *App) SetMeta(meta Meta) *App {
	app.meta = meta
	return app
}

var userInputCmd string = ""
var userInputArgs []string = []string{}

func (app *App) Run() {
	if len(os.Args) < 2 {
		fmt.Println("No command specified")
		return
	}

	userInputCmd = os.Args[1]
	userInputArgs = os.Args[2:]

	for _, command := range app.commands {
		if command.Key == userInputCmd {
			command.Call(*app)
			return
		}
	}

	fmt.Println("Unknown command:", userInputCmd)
}

func (app *App) GetMeta() Meta {
	return app.meta
}

func (app *App) Command(key string, call func(App)) *command {
	app.commands = append(app.commands, command{
		Call: call,
		Key:  key,
	})
	return &app.commands[len(app.commands)-1]
}

func (command *command) SetMeta(meta Meta) {
	command.Meta = meta
}

func (app *App) HasArgs(name string) bool {
	return slices.Contains(userInputArgs, name)
}

func (app *App) GetArgsByIndex(index int) string {
	if len(userInputArgs) > index {
		return userInputArgs[index]
	}
	return ""
}

func (app *App) GetArgsByName(name string) string {
	if index := slices.Index(userInputArgs, name); index >= 0 {
		if len(userInputArgs)-1 >= index+1 {
			return userInputArgs[index+1]
		}
	}
	return ""
}

func (app *App) GetArgsByEqual(name string) string {
	for _, arg := range userInputArgs {
		_arg := strings.Split(arg, "=")
		if len(_arg) == 2 && _arg[0] == name {
			return _arg[1]
		}
	}
	return ""
}

func Help(app App) {
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
