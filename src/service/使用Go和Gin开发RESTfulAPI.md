# 使用GO和GIN开发Restful API

------

## 创建一个模块

```
go mod init example/web-service-gin
```

创建一个web-service-gin 模块



## 创建数据



创建一个 main..go 文件

在 main.go 文件的顶部，粘贴以下包声明。

```
package main
```

创建一个音乐专辑的结构

```go
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
```

创建这个结构下的专辑数据

```go
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```



## 编写一个处理程序来返回所有专辑

```
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}
```

**建立 getAlbums 函数 与  `/albums` 断端点的联系**

```go
func main() {
    router := gin.Default()					// 创建一个默认的路由器
    router.GET("/albums", getAlbums)		// 使用 GET 函数将 GET HTTP 方法和 /albums 路径与处理程序函数相关联。

    router.Run("localhost:8080")			// 使用 Run 函数将路由器连接到 http.Server 并启动服务器。
}
```

导入需要的包

```go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)
```

测试 API

```shell
go run .
curl http://localhost:8080/albums
```

```json
[
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        }
]
```



## 编写一个处理程序来添加一个新专辑

```go
// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
```

在main 函数中绑定

```go
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)			// GET 方法 获取所以有
    router.POST("/albums", postAlbums)			// POST 方法添加

    router.Run("localhost:8080")
}
```

```json
POST	http://localhost:8080/albums

{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}

GET 	http://localhost:8080/albums

[
    {
        "id": "1",
        "title": "Blue Train",
        "artist": "John Coltrane",
        "price": 56.99
    },
    {
        "id": "2",
        "title": "Jeru",
        "artist": "Gerry Mulligan",
        "price": 17.99
    },
    {
        "id": "3",
        "title": "Sarah Vaughan and Clifford Brown",
        "artist": "Sarah Vaughan",
        "price": 39.99
    },
    {
        "id": "4",
        "title": "The Modern Sound of Betty Carter",
        "artist": "Betty Carter",
        "price": 49.99
    }
]
```

## 编写处理程序以返回特定专辑

```go

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```

在main 函数中绑定

```go
func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
```



```json
http://localhost:8080/albums/1
{
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
}

http://localhost:8080/albums/2
{
    "id": "2",
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
}
```

