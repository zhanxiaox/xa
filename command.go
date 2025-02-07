package xa

type command struct {
	name        string
	description string
	arguments   []argument
	call        func(*App)
}

type argument struct {
	name        string
	description string
}

func (c *command) Desc(desc string) *command {
	c.description = desc
	return c
}

func (c *command) Args(name, desc string) *command {
	c.arguments = append(c.arguments, argument{name: name, description: desc})
	return c
}

func (c *command) GetName() string {
	return c.name
}

func (c *command) GetDescription() string {
	return c.description
}

func (c *command) GetArgments() []argument {
	return c.arguments
}

func (a *argument) GetName() string {
	return a.name
}

func (a *argument) GetDescription() string {
	return a.description
}
