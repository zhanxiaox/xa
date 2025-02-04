package cmds

import (
	"fmt"

	"github.com/zhanxiaox/xa"
)

func Install(app *xa.App) {
	fmt.Println(app.GetRuntime().Cmd)
}

func Uninstall(app *xa.App) {
	fmt.Println(app.GetRuntime().Cmd)
}
