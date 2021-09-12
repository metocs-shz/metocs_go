# 															GO语言基础

------

## 安装GO

​	下载地址  https://golang.google.cn/dl/

------

## 第一个GO程序

一个简短的 Hello, World 入门教程。  了解一些关于 Go 代码、工具、包和模块的知识。

### Prerequisites ——先决条件

- ​    **Some programming experience.** The code here is pretty    simple, but it helps to know something about functions.  
- ​    **A tool to edit your code.** Any text editor you have will    work fine. Most text editors have good support for Go. The most popular are    VSCode (free), GoLand (paid), and Vim (free).  
- ​    **A command terminal.** Go works well using any terminal on    Linux and Mac, and on PowerShell or cmd in Windows.  



配置 go环境变量

%GO_ROOT%\bin 

cmd 测试 go

**go version**

```shell
go version go1.17.1 windows/amd64
```

**Hello world**

创建一个 hello.go 文件

```go
package main		//声明一个 main包（包是一种分组方式 函数，它由同一目录中的所有文件组成）。

import "fmt"		//导入通用的 fmt包， 这是一个格式化文本包 包含打印到控制台的功能.

func main(){		// 默认执行的main 函数 会打印 Hello 到控制台.	
	imt.Println("Hello World!")
}
```

为代码所在模块命名

```
D:\demo\metocs_go\src>go mod init helloGo
go: creating new go.mod: module helloGo
go: to add module requirements and sums:
        go mod tidy

D:\demo\metocs_go\src>
```

执行 命令   go  run  [文件名]

```shell
D:\demo\metocs_go\src>go run hello.go
Hello, World!

D:\demo\metocs_go\src>
```



### 调用外部包中的代码

创建文件  quote.go 

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

先了解 go mod 命令  以了解关于modules

```shell
D:\demo\metocs_go\src>go mod help
Go mod provides access to operations on modules.

Note that support for modules is built into all the go commands,
not just 'go mod'. For example, day-to-day adding, removing, upgrading,
and downgrading of dependencies should be done using 'go get'.
See 'go help modules' for an overview of module functionality.

Usage:

        go mod <command> [arguments]

The commands are:

        download    download modules to local cache				#下载模块到本地缓存 
        edit        edit go.mod from tools or scripts			#从工具或脚本编辑 go.mod 
        graph       print module requirement graph				#打印模块需求图 
        init        initialize new module in current directory	#在当前目录初始化新模块 
        tidy        add missing and remove unused modules		#添加缺失的并删除未使用的模块 
        vendor      make vendored copy of dependencies			#制作依赖项的供应商副本 
        verify      verify dependencies have expected content	#验证依赖项是否具有预期内容 
        why         explain why packages or modules are needed	#解释为什么需要包或模块 

Use "go help mod <command>" for more information about a command.


D:\demo\metocs_go\src>
```

 下载缺少的模块 删除未使用的模块

```shell
D:\demo\metocs_go\src>go mod tidy					
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c

D:\demo\metocs_go\src>
```

运行程序 执行 quote.go 文件

```
D:\demo\metocs_go\src>go run quote.go
Don't communicate by sharing memory, share memory by communicating.

D:\demo\metocs_go\src>
```

------

## 创建一个GO模块

### 创建一个可以被调用 并返回结果的程序

为您的Go 模块源代码创建一个 greetings 目录

```shell
mkdir greetings
cd greetings
```

创建模块

```shell
D:\demo\metocs_go\src\greetings>go mod init example.com/greetings
go: creating new go.mod: module example.com/greetings

D:\demo\metocs_go\src\greetings>
```

创建 greetings.go 程序文件

```go
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)                 //   在 Go 中， :=  运算符是声明和的快捷方式 在一行中初始化一个变量
    return message
}
```

![image-20210911231819568](C:\Users\xiaomy\AppData\Roaming\Typora\typora-user-images\image-20210911231819568.png)



### 创建一个调用其他模块的程序

为您的 Go 模块源代码创建一个 hello 目录。  这是调用程序的目录。

创建模块

```shell
D:\demo\metocs_go\src\hello>go mod init example.com/hello
go: creating new go.mod: module example.com/hello

D:\demo\metocs_go\src\hello>
```



创建  hello.go.  程序文件

```go
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("metocs")
    fmt.Println(message)
}
```



请使用   go mod edit  命令编辑 `example.com/hello`模块从其模块路径（模块不在的位置）重定向 Go 工具 到本地目录（它所在的位置）。

在 hello 目录的命令提示符下，运行以下命令         

命令：

```go
go mod edit -replace example.com/greetings=../greetings
```

该命令会在  go.mod 文件中加入

```go
replace example.com/greetings => ../greetings

变为

module example.com/hello

go 1.17

replace example.com/greetings => ../greetings
```

在 hello 目录的命令提示符下，运行     go mod tidy  命令同步 example.com/hello  模块的依赖项，添加那些 代码需要，但尚未在模块中module。

```
go mod tidy


会在 go.mod  中加入

require example.com/greetings v0.0.0-00010101000000-000000000000


可以修改版本

require example.com/greetings v1.1.0

```



运行调用程序代码

```go
go run .

D:\demo\metocs_go\src\hello>go run .
Hi, metocs. Welcome!

D:\demo\metocs_go\src\hello>
```









