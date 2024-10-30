package xa

import (
	"fmt"
	"os"
)

type runtime struct {
	ExcuteCommand     string
	ExcuteCommandName string
	ExcuteArgs        map[string]any
}

func (r *runtime) Excute() {
	fmt.Println(os.Args[0], os.Args[1])
	r.ExcuteCommandName = os.Args[1]
}

var rt runtime

func New() *app {
	return &app{cmds: map[string]*Command{}}
}

func (a *app) Run() {
	rt.Excute()
	if cmd, ok := a.cmds[rt.ExcuteCommandName]; ok {
		fmt.Println("988888-", cmd.call, cmd.desc, cmd.name, cmd.args)

	}
}

func (a *app) Name(name string) *app {
	a.name = name
	return a
}

func (a *app) Desc(desc string) *app {
	a.desc = desc
	return a
}

func (a *app) Version(ver string) *app {
	a.version = ver
	return a
}

func (a *app) Author(author string) *app {
	a.author = author
	return a
}

func (a *app) Load(cmd, desc string) *Command {
	c := Command{name: cmd, desc: desc}
	a.cmds[cmd] = &c
	return &c
}

type app struct {
	name    string
	desc    string
	author  string
	version string
	cmds    map[string]*Command
}

type Command struct {
	name string
	desc string
	args []Args
	call func(*Command)
}

type Args struct {
	name string
	desc string
}

func (c *Command) Desc(desc string) *Command {
	c.desc = desc
	return c
}

func (c *Command) Args(key, desc string) *Command {
	c.args = append(c.args, Args{name: key, desc: desc})
	return c
}
func (c *Command) Call(fn func(*Command)) *Command {
	fmt.Println("call1:", c)
	c.call = fn
	fmt.Println("call2:", c)
	return c
}

func (c *Command) GetName() string {
	return c.name
}

func (c *Command) GetDesc() string {
	return c.desc
}

func (c *Command) GetCall() func(*Command) {
	return c.call
}

func (c *Command) GetArgs() []Args {
	return c.args
}

func (c *Command) Get(name string) string {
	return "a"
}
