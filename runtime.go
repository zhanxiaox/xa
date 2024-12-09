package xa

import (
	"os"
	"slices"
)

type Runtime struct {
	path string
	name string
	args []string
}

func (r *Runtime) Execute() {
	r.path = os.Args[0]
	if len(os.Args) > 1 {
		r.name = os.Args[1]
	}
	if len(os.Args) > 2 {
		r.args = os.Args[2:]
	}
}

var rt Runtime

func (r *Runtime) GetPath() string {
	return r.path
}

func (r *Runtime) GetName() string {
	return r.name
}

func (r *Runtime) GetArg(arg string) string {
	index := slices.Index(r.args, arg)
	if index > 0 && len(os.Args) >= index+1 {
		return os.Args[index+1]
	}
	return ""
}

func (r *Runtime) HasArg(arg string) bool {
	return slices.Contains(r.args, arg)
}
