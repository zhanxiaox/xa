package cmds

import (
	"fmt"
	"os"

	"gitee.com/zhanxiaox/xa"
)

func Install(app xa.App) {
	fmt.Println(os.Args[0])
}

func Uninstall(app xa.App) {
	fmt.Println(os.Args[0])
}
