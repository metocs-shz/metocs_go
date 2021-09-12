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
	fmt.Println("Hello World!")
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

### 创建一个可以被调用返回结果的程序

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



### 创建一个可以返回并处理错误的程序

修改 greetings 模块代码

```go
package greetings

import (

	"fmt"
	//  导入 Go 标准库错误包，以便您可以使用它的 errors.New 函数。
	"errors"
)

func Hello(name string) (string,error) {		//更改函数，使其返回两个值：一个字符串和一个错误。  您的调用者将检查第二个值以查看是否发生错误。
	
	 // 添加 if 语句以检查无效请求（名称应为空字符串）并在请求无效时返回错误。  errors.New 函数返回一个错误，其中包含您的消息。
    if name == "" {
        return "", errors.New("empty name")		
    }
    
    message := fmt.Sprintf("Hi, %v. Welcome!", name)								//  传递了就返回一个嵌入名字的值
    return message, nil							// 添加 nil（意味着没有错误）作为成功返回的第二个值。  这样，调用者就可以看到函数成功了。
}
```

 **<u>（任何 Go 函数都可以返回多个值。有关更多信息，请参阅 Effective Go。）</u>**



**修改调用者程序  hello.go**

```go
package main

import (

	"fmt"

	"log"						// 导入 log 包以打印日志

	"example.com/greetings"
)

func main(){
    log.SetPrefix("greetings: ")		// 设置日志开头
    log.SetFlags(0)
    
    message, err := greetings.Hello("")							// 获取被调用程序返回结果  传递 空字符串
   
    if err != nil {												// 判断错误是否为空
        log.Fatal(err)											// 打印错误信息
    }

    fmt.Println(message)										// 输出不为空的返回值到控制台
}
```



**测试** 			执行 hello.go程序

```
D:\demo\metocs_go\src\hello>go run .
greetings: empty name
exit status 1

D:\demo\metocs_go\src\hello>
```



### **留坑 errors 包 与 log 包**

### log 包

**概述**

包日志实现了一个简单的日志包。  它定义了一个类型，Logger， 与格式化输出的方法。  它还具有预定义的“标准” 可通过辅助函数 Print[f|ln]、Fatal[f|ln] 和 Panic[f|ln]，比手动创建 Logger 更容易使用。 该记录器写入标准错误并打印日期和时间 每个记录的消息。 每条日志消息都在单独的行上输出：如果消息是 打印不以换行符结尾，记录器将添加一个。 Fatal 函数在写入日志消息后调用 os.Exit(1)。 Panic 函数在写入日志消息后调用 panic。 

**Constants**

```go
const (
	Ldate = 1 << iota      								// 本地时区的日期：2009/01/23 
	Ltime 												// 本地时区的时间 
	Lmicroseconds 										// 微秒分辨率：01:23:23.123123。 假设 Ltime。 
	Llongfile 											// 完整文件名和行号：/a/b/c/d.go:23 
	Lshortfile 											// 最终文件名元素和行号：d.go:23。 覆盖 Llongfile 
	LUTC 												// 如果设置了 Ldate 或 Ltime，则使用 UTC 而不是本地时区 
	Lmsgprefix 											// 将“前缀”从行首移动到消息之前 
	LstdFlags = Ldate | Ltime 							// 标准记录器的初始值
)
```

这些常量的组合会修改打印内容

flags Ldate | Ltime (or LstdFlags) produce,

```
2009/01/23 01:23:23 message
```

while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,

```
2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
```

**Functions** 

func fatal

```go
func Fatal(v ...interface{})

Fatal 相当于 Print() 然后再调用 os.Exit(1)
```

```go
package main
import (
//    "fmt"
    "errors"
    "log"
)
func main() {
   // fmt.Println("Hello, World!")
   message := errors.New("------->测试 Fatal")
   log.Fatal(message)
}

D:\demo\metocs_go\src\helloGo>go run hello.go
2021/09/12 10:50:32 ------->测试 Fatal
exit status 1

D:\demo\metocs_go\src\helloGo>
```

**注意**

```shell
D:\demo\metocs_go\src\helloGo>go run hello.go				# 导入没有用到的包编译会报错
# command-line-arguments
.\hello.go:4:5: imported and not used: "fmt"

D:\demo\metocs_go\src\helloGo>  
```

func Fatalf

```go
func Fatalf(format string, v ...interface{})

Fatalf 相当于 Printf() 然后再调用 os.Exit(1)
```

```go
package main

import (
   // "fmt"
    "errors"
    "log"
)
func main() {
   // fmt.Println("Hello, World!")
   message := errors.New("------->测试 Fatalf")
   log.Printf("就是测试这个  %s, %s ",message,"对就是这个")
   log.Fatalf("就是测试这个  %s, %s ",message,"对就是这个")
}

D:\demo\metocs_go\src\helloGo>go run hello.go
2021/09/12 10:58:04 就是测试这个  ------->测试 Fatalf, 对就是这个
2021/09/12 10:58:04 就是测试这个  ------->测试 Fatalf, 对就是这个
exit status 1

D:\demo\metocs_go\src\helloGo>
```

func Fatalln

```go
func Fatalln(v ...interface{})

Fatalln 相当于 Println() 然后再调用 os.Exit(1)
```

```go
package main
import (
    //"fmt"
    "errors"
    "log"
)
func main() {
   // fmt.Println("Hello, World!")
   message := errors.New("------->测试 Fatal")

   log.Println(message)
   log.Fatalln("Hello World!")

}

D:\demo\metocs_go\src\helloGo>go run hello.go
2021/09/12 10:56:10 ------->测试 Fatal
2021/09/12 10:56:10 Hello World!
exit status 1

D:\demo\metocs_go\src\helloGo>
```



**注意**

**Fatal、Fatalf、Fatalln 等函数调用后程序就会退出 不会执行之后的代码**

func Flags

```go
func Flags() int 

Flags 返回标准记录器的输出标志。 标志位是 Ldate、Ltime 等。 
```

```go
package main
import (
    //"fmt"
    "errors"
    "log"
)
func main() {
   // fmt.Println("Hello, World!")
   message := errors.New("------->测试 Fatal")
   log.Println(message)
   num := log.Flags()
   log.Fatalln(num)
}

D:\demo\metocs_go\src\helloGo>go run hello.go
2021/09/12 11:10:54 ------->测试 Fatal
2021/09/12 11:10:54 3
exit status 1

D:\demo\metocs_go\src\helloGo>
```

func Output

```go
func Output(calldepth int, s string) error

输出写入日志事件的输出。 字符串 s 包含 在由标志指定的前缀之后打印的文本记录。 如果 s 的最后一个字符不是，则附加换行符已经换行了。 Calldepth 是数量的计数 计算文件名和行号时要跳过的帧 如果设置了 Llongfile 或 Lshortfile；值为 1 将打印详细信息对于输出的调用者。
```

func SetPrefix

```go
func SetPrefix(prefix string)

设置logger的输出前缀。
```

```go
package main
import (
    //"fmt"
    "errors"
    "log"
)
func main() {
   log.SetFlags(log.Ldate |log.Ltime | log.Lmicroseconds )	
   log.SetPrefix("看看是啥----->")
   message := errors.New("--------->  错误！！！")
   log.Print(message)
}

D:\demo\metocs_go\src\helloGo>go run hello.go
看看是啥----->2021/09/12 11:35:13.444140 --------->  错误！！！

D:\demo\metocs_go\src\helloGo>
```



**后续会在进行整理改正**

### 返回水随机问候

修改greeting.go 程序

```go
package greetings
import (
    "errors"
    "fmt"
    "math/rand"												// 用于生成随机数
    "time"													// 用于获取时间
)
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// init sets initial values for variables used in the function.			// init 方法 用于在程序加载时执行   一个程序文件可以有有多个 init 方法
func init() {
    rand.Seed(time.Now().UnixNano())
}

func init() {
    fmt.println("我已经执行完了")
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
```

```
D:\demo\metocs_go\src\hello>go run .
我已经执行完了
greetings: empty name
exit status 1

D:\demo\metocs_go\src\hello>
```

```
D:\demo\metocs_go\src\hello>go run .
我已经执行完了
Great to see you, metocs!

D:\demo\metocs_go\src\hello>go run .
我已经执行完了
Hail, metocs! Well met!

D:\demo\metocs_go\src\hello>go run .
我已经执行完了
Hi, metocs. Welcome!

D:\demo\metocs_go\src\hello>
```























