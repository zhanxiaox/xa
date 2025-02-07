package xa

import (
	"errors"
	"fmt"
	"os"
	"slices"
)

type App struct {
	info     AppInfo
	commands map[string]*command
	runtime  runtime
}

func New(info AppInfo) *App {
	return &App{
		info:     info,
		commands: make(map[string]*command),
		runtime:  runtime{},
	}
}

type AppInfo struct {
	Name              string
	Description       string
	Author            string
	Contact           string
	Version           string
	Usage             string
	EnableDefaultHelp bool
}

type runtime struct {
	path      string
	command   string
	arguments []string
}

func (app *App) GetAppInfo() AppInfo {
	return app.info
}

func (app *App) GetRuntime() (string, string, []string) {
	return app.runtime.path, app.runtime.command, app.runtime.arguments
}

func (app *App) HasArgument(argument string) bool {
	return slices.Contains(app.runtime.arguments, argument)
}

func (app *App) GetArgument(argument string) string {
	index := slices.Index(app.runtime.arguments, argument)
	if index > 0 && len(os.Args) >= index+1 {
		return app.runtime.arguments[index+1]
	}
	return ""
}

func (app *App) Run() {
	app.runtime.path = os.Args[0]
	if len(os.Args) > 1 {
		app.runtime.command = os.Args[1]
	}
	if len(os.Args) > 2 {
		app.runtime.arguments = os.Args[2:]
	}
	if app.info.EnableDefaultHelp {
		app.Command("help", defaultHelp).Desc("Print this default help information")
	}
	if command, ok := app.commands[app.runtime.command]; ok {
		command.call(app)
	} else {
		if app.info.EnableDefaultHelp {
			defaultHelp(app)
		} else {
			fmt.Println("Unknown command:", app.runtime.command)
		}
	}
}

func (app *App) Command(name string, call func(*App)) (c *command) {
	c = &command{name: name, call: call}
	app.commands[name] = c
	return
}

func (app *App) GetCommand(name string) (command, error) {
	if cmd, ok := app.commands[name]; !ok {
		return command{}, errors.New("command not found")
	} else {
		return *cmd, nil
	}
}

func defaultHelp(app *App) {
	fmt.Println(app.info.Name, app.info.Version)
	fmt.Println(app.info.Description)
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println(app.info.Usage)
	fmt.Println()
	fmt.Println("OPTIONS:")
	for _, cmd := range app.commands {
		fmt.Printf("%-17s %v", cmd.name, cmd.description)
		for _, argument := range cmd.arguments {
			fmt.Printf("\n  %-15s %v", argument.name, argument.description)
		}
		fmt.Println()
	}
}
