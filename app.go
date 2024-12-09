package xa

type app struct {
	name    string
	desc    string
	author  string
	contact string
	version string
	cmds    map[string]*Command
}

func (a *app) Name(name string) *app {
	a.name = name
	return a
}

func (a *app) GetName() string {
	return a.name
}

func (a *app) Desc(desc string) *app {
	a.desc = desc
	return a
}

func (a *app) GetDesc() string {
	return a.desc
}

func (a *app) Version(ver string) *app {
	a.version = ver
	return a
}

func (a *app) GetVersion() string {
	return a.version
}

func (a *app) Author(author string) *app {
	a.author = author
	return a
}

func (a *app) GetAuthor() string {
	return a.author
}

func (a *app) Run() {
	rt.Execute()
	if cmd, ok := a.cmds[rt.GetName()]; ok {
		cmd.call(GetRuntime())
	}
}

func (a *app) Load(cmd string, call func(*Runtime)) *Command {
	c := Command{name: cmd, call: call}
	if a.cmds == nil {
		a.cmds = map[string]*Command{}
	}
	a.cmds[cmd] = &c
	return &c
}

var a app
