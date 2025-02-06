package xa

type command struct {
	Name        string
	Description string
	Argments    []args
	call        func(*App)
}

type args struct {
	Name        string
	Description string
}

func (c *command) Desc(desc string) *command {
	c.Description = desc
	return c
}

func (c *command) Args(name, desc string) *command {
	c.Argments = append(c.Argments, args{Name: name, Description: desc})
	return c
}
