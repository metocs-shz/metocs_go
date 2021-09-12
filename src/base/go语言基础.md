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

**后续会在进行整理改正**

### 返回随机问候

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

### 回复多人问候  

在greeting.go 程序中创建新的函数

```go
func Hellos(names []string) (map[string]string,error){
    messages := make(map[string]string)						// 创建一个map key sring name value string messgae
    														// map 初始化的语法   make(map[key-type]value-type)
    for _ , name := range names {							// 在这个for循环中range将会返回两个值：索引 与 数组中的值的副本
        													// 颗不需要哪个值可以用空白标识符代替
        message ,err := Hello(name)			
        if err != nil {
            return nil,err
        }
        messages[name] = messages
    }
    return messages,nil
}
```

注意  该 for循环中返回的为数组中值的副本



修改 hello.go 文件

```go
package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)
	
func main(){
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    //message, err := greetings.Hello("metocs")

    names := []string{"Gladys", "Samantha", "Darrin","metocs"}		// 创建一个字符串数组

    messages, err := greetings.Hellos(names)						// 调用Hellos 方法

    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)											// 打印输出结果
}
```



```
D:\demo\metocs_go\src\hello>go run .
map[Darrin:Great to see you, Darrin! Gladys:Great to see you, Gladys! Samantha:Hail, Samantha! Well met! metocs:Hi, metocs. Welcome!]

D:\demo\metocs_go\src\hello>
```



### 创建一个测试单元

Go 对单元测试的内置支持使您可以更轻松地进行测试。   具体来说，使用命名约定，Go 的 `testing`包，和   这 `go test`命令，您可以快速编写和执行测试。 

1. ​    在 greetings 目录中，创建一个名为 greetings_test.go 的文件。      

   ​      以 _test.go 结尾的文件名告诉 `go test`命令  该文件包含测试功能      

2. 复制以下代码

3. 

```go
package greetings

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}


/*
    创建两个测试函数来测试 greetings.Hello功能。 测试函数名称具有以下形式 TestName, 其中 Name 说明了有关特定测试的内容。 还有，测试函数接受一个指向 testing包的 testing.T类型 作为参数。 使用此参数的方法进行报告并从您的测试中记录。 
*/
```

使用   go test   命令进行测试

```
D:\demo\metocs_go\src\greetings>go test
PASS
ok      example.com/greetings   0.046s

D:\demo\metocs_go\src\greetings>
```

使用   go test   -v    命令进行测试  并获得更详细的结果

```
D:\demo\metocs_go\src\greetings>go test -v
=== RUN   TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN   TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok      example.com/greetings   0.036s

D:\demo\metocs_go\src\greetings>
```



修改 greetings.go  Hello 函数代码

```go
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return name, errors.New("empty name")
    }
    // Create a message using a random format.
    // message := fmt.Sprintf(randomFormat(), name)			
    message := fmt.Sprint(randomFormat())				// 即使传递了名字也不会将名字放入
    return message, nil
}

//因测试代码中

func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)			// 对返回结果是否将该名字放入其中进行了判断 因此 
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)		// 这一行代码被执行了
    }
}


```

### 编译并安装应用程序 

  在最后一个主题中，您将学习一些新的 `go`命令。  尽管 `go run`命令是当您进行频繁更改程序时快捷编译和运行的方式 ，它不会生成可执行的二进制文件 

  本主题介绍了两个用于构建代码的附加命令： 

- ​    这 go   build  命令 编译包及其依赖项，但它不会安装编译结果。    
- ​    这 go  install  命令 编译并安装软件包。   

在 hello 模块中 执行 以下命令

```
go build
```

![image-20210912180311168](C:\Users\xiaomy\AppData\Roaming\Typora\typora-user-images\image-20210912180311168.png)

获得以上文件 并在当前目录下执行该文件

```
D:\demo\metocs_go\src\hello>hello.exe
map[Darrin:Hail, Darrin! Well met! Gladys:Hi, Gladys. Welcome! Samantha:Hi, Samantha. Welcome! metocs:Hail, metocs! Well met!]

D:\demo\metocs_go\src\hello>
```

发现 Go 安装路径，go 命令将在其中安装当前包。

```
go list -f '{{.Target}}'
```

例如，命令的输出可能会说 `/home/gopher/bin/hello`,       意味着二进制文件安装到 /home/gopher/bin。  你会需要这个，   下一步安装目录。     

将 Go 安装目录添加到系统的 shell 路径。      

```
 set PATH=%PATH%;D:\software\go				// 设置 go的安目录
 
 
 go env -w GOBIN=D:\demo\metocs_go\bin		// 设置go install 后文件的存放目录 
```

hello 下执行 go install



hello.exe 就会放在  D:\demo\metocs_go\bin 目录下  在任何文件夹下访问 该文件都可以的到结果





