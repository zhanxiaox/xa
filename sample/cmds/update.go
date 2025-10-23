package cmds

import (
	"fmt"

	"gitee.com/zhanxiaox/xa"
)

func Update(app xa.App) {
	fmt.Println("-y", app.HasArgs("-y"))
	fmt.Println("-v", app.HasArgs("-v"))
	fmt.Println("-f", app.HasArgs("-f"))
}
