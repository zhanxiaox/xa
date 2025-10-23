package cmds

import (
	"fmt"

	"gitee.com/zhanxiaox/xa"
)

func Version(app xa.App) {
	fmt.Println(app.GetMeta().Version)
}
