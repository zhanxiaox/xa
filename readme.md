# xa

**xa** 是一个轻量级的Go语言命令行应用框架，帮助开发者快速构建功能丰富的命令行工具。

## 特性

- 简洁的命令注册机制
- 灵活的命令行参数解析
- 自动生成帮助信息
- 支持子命令和参数
- 支持多种参数格式（位置参数、命名参数、键值参数）

## 安装

使用Go模块安装：

```bash
go get gitee.com/zhanxiaox/xa
```

## 快速开始

下面是一个简单的示例，展示如何使用xa创建一个命令行应用：

```go
package main

import (
	"fmt"
	"gitee.com/zhanxiaox/xa"
)

func main() {
	// 创建应用实例
	app := xa.New(xa.Meta{
		Name:        "myapp",
		Author:      "Your Name",
		Version:     "1.0.0",
		Description: "My awesome command line application",
		Usage:       "myapp [command] [options]",
		Contact:     "your@email.com",
	})

	// 注册命令
	app.Command("hello", xa.Meta{
		Call:        helloCommand,
		Description: "Say hello to someone",
		Params: []xa.Meta{
			{Name: "--name", Description: "Name to greet"},
		},
	})

	// 运行应用
	app.Run()
}

// 命令处理函数
func helloCommand(app xa.App) {
	name := app.GetArgsByName("--name")
	if name == "" {
		name = "World"
	}
	fmt.Printf("Hello, %s!\n", name)
}
```

## 使用指南

### 创建应用

使用`New()`函数创建一个新的应用实例：

```go
app := xa.New(xa.Meta{
	Name:        "appname",      // 应用名称
	Author:      "Author Name",  // 作者名称
	Version:     "1.0.0",        // 版本号
	Description: "App description", // 应用描述
	Usage:       "appname [command]", // 使用说明
	Contact:     "contact@email", // 联系方式
})
```

### 注册命令

使用`Command()`方法注册命令：

```go
app.Command("commandName", xa.Meta{
	Call:        commandHandler,  // 命令处理函数
	Description: "Command description", // 命令描述
	Params: []xa.Meta{  // 命令参数描述
		{Name: "--param1", Description: "Parameter 1 description"},
		{Name: "--param2", Description: "Parameter 2 description"},
	},
})
```

### 运行应用

使用`Run()`方法启动应用：

```go
app.Run()
```

这将解析命令行参数并执行相应的命令处理函数。

### 参数解析

xa提供了多种方法来获取命令行参数：

#### 按索引获取参数

```go
arg := app.GetArgsByIndex(0) // 获取第一个参数
```

#### 按名称获取参数值

```go
value := app.GetArgsByName("--name") // 获取 --name 后面的值
```

#### 检查参数是否存在

```go
hasHelp := app.HasArgs("--help") // 检查 --help 参数是否存在
```

#### 按键值格式获取参数

```go
value := app.GetArgsByEqual("name") // 获取 name=value 格式的值
```

### 内置帮助命令

xa提供了内置的`Help`函数，可以直接用作帮助命令的处理函数：

```go
app.Command("help", xa.Meta{
	Call:        xa.Help,
	Description: "Show help information",
})
```

## 高级用法

### 多级命令

虽然xa核心不直接支持嵌套的多级命令，但你可以通过在命令处理函数中实现子命令逻辑来实现类似功能：

```go
func parentCommand(app xa.App) {
	if len(app.GetArgsByIndex(0)) > 0 {
		childCmd := app.GetArgsByIndex(0)
		switch childCmd {
		case "sub1":
			subCommand1(app)
		case "sub2":
			subCommand2(app)
		default:
			fmt.Println("Unknown subcommand:", childCmd)
		}
	} else {
		fmt.Println("Parent command requires a subcommand")
	}
}
```

## 项目结构

```
xa/
├── app.go       # 核心代码
├── go.mod       # Go模块文件
├── .gitignore   # Git忽略文件
└── readme.md    # 项目文档
```

## 示例

可以参考项目中的sample目录，其中包含了一个完整的使用示例。

## 许可证

[MIT](LICENSE)

## 贡献

欢迎提交Issue和Pull Request！

## 作者

zhanxiaox - 346084070@qq.com