package xa

type Command struct {
	name string
	desc string
	args []Args
	call func(*App)
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
func (c *Command) Call(fn func(*App)) *Command {
	c.call = fn
	return c
}

func (c *Command) GetName() string {
	return c.name
}

func (c *Command) GetDesc() string {
	return c.desc
}

func (c *Command) GetCall() func(*App) {
	return c.call
}

func (c *Command) GetArgs() []Args {
	return c.args
}
