package cmds

import (
	"fmt"

	"github.com/zhanxiaox/xa"
)

func Version(app *xa.App) {
	fmt.Println(app.GetAppInfo().Version)
}
