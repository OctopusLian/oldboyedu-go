<!--
 * @Description: 
 * @Author: neozhang
 * @Date: 2022-04-05 14:44:14
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-05 15:28:04
-->
# gin框架  

## gin入门  

### 介绍  

- Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点  
- 对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错  
- 借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范  

### 安装  

`import github.com/gin-gonic/gin`  

### HelloWorld  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

// gin的helloWorld

func main() {
   // 1.创建路由
   r := gin.Default()
   // 2.绑定路由规则，执行的函数
   // gin.Context，封装了request和response
   r.GET("/", func(c *gin.Context) {
      c.String(http.StatusOK, "hello World!")
   })
   // 3.监听端口，默认在8080
   r.Run(":8000")
}
```

## gin路由  

### 基本路由  

- gin 框架中采用的路由库是基于httprouter做的  
- 地址为：<https://github.com/julienschmidt/httprouter>  

### Restful风格的API  

- gin支持Restful风格的AP  
- 即Representational State Transfer的缩写。直接翻译的意思是"表现层状态转化"，是一种互联网应用程序的API设计理念：URL定位资源，用HTTP描述操作  

1.获取文章 /blog/getXxx      Get   blog/Xxx  
2.添加     /blog/addXxx      POST  blog/Xxx  
3.修改     /blog/updateXxx   PUT   blog/Xxx  
4.删除     /blog/delXxxx     DELETE blog/Xxx  

### API参数  

- 可以通过Context的Param方法来获取API参数  
- localhost:8000/xxx/zhangsan  

### URL参数  

- URL参数可以通过DefaultQuery()或Query()方法获取  
- DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串  
- API ? name=zs  

### 表单参数  

- 表单传输为post请求，http常见的传输格式为四种：application/json,application/x-www-form-urlencoded,application/xml,multipart/form-data  
- 表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数  

### 上传单个文件  

- multipart/form-data格式用于文件上传  
- gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中  

### 上传多个文件  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
   "fmt"
)

// gin的helloWorld

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // 限制表单上传大小 8MB，默认为32MB
   r.MaxMultipartMemory = 8 << 20
   r.POST("/upload", func(c *gin.Context) {
      form, err := c.MultipartForm()
      if err != nil {
         c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
      }
      // 获取所有图片
      files := form.File["files"]
      // 遍历所有图片
      for _, file := range files {
         // 逐个存
         if err := c.SaveUploadedFile(file, file.Filename); err != nil {
            c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
            return
         }
      }
      c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
   })
   r.Run(":8000")
}
```

### routes group  

- routes group是为了管理一些相同的URL  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "fmt"
)

// gin的helloWorld

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // 路由组1 ，处理GET请求
   v1 := r.Group("/v1")
   // {} 是书写规范
   {
      v1.GET("/login", login)
      v1.GET("submit", submit)
   }
   v2 := r.Group("/v2")
   {
      v2.POST("/login", login)
      v2.POST("/submit", submit)
   }
   r.Run(":8000")
}

func login(c *gin.Context) {
   name := c.DefaultQuery("name", "jack")
   c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
   name := c.DefaultQuery("name", "lily")
   c.String(200, fmt.Sprintf("hello %s\n", name))
}
```

### 路由原理  

- httproter会将所有路由规则构造一颗前缀树  
- 例如有 root and as at cn com  

## gin数据解析和绑定  

### json数据解析和绑定  

客户端传参，后端接收并解析到结构体  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

// 定义接收数据的结构体
type Login struct {
   // binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
   User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
   Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // JSON绑定
   r.POST("loginJSON", func(c *gin.Context) {
      // 声明接收的变量
      var json Login
      // 将request的body中的数据，自动按照json格式解析到结构体
      if err := c.ShouldBindJSON(&json); err != nil {
         // 返回错误信息
         // gin.H封装了生成json数据的工具
         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
         return
      }
      // 判断用户名密码是否正确
      if json.User != "root" || json.Pssword != "admin" {
         c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status": "200"})
   })
   r.Run(":8000")
}
```

### 表单数据解析和绑定  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

// 定义接收数据的结构体
type Login struct {
   // binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
   User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
   Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // JSON绑定
   r.POST("/loginForm", func(c *gin.Context) {
      // 声明接收的变量
      var form Login
      // Bind()默认解析并绑定form格式
      // 根据请求头中content-type自动推断
      if err := c.Bind(&form); err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
         return
      }
      // 判断用户名密码是否正确
      if form.User != "root" || form.Pssword != "admin" {
         c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status": "200"})
   })
   r.Run(":8000")
}
```

### URI数据解析和绑定  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

// 定义接收数据的结构体
type Login struct {
   // binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
   User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
   Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // JSON绑定
   r.GET("/:user/:password", func(c *gin.Context) {
      // 声明接收的变量
      var login Login
      // Bind()默认解析并绑定form格式
      // 根据请求头中content-type自动推断
      if err := c.ShouldBindUri(&login); err != nil {
         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
         return
      }
      // 判断用户名密码是否正确
      if login.User != "root" || login.Pssword != "admin" {
         c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
         return
      }
      c.JSON(http.StatusOK, gin.H{"status": "200"})
   })
   r.Run(":8000")
}
```

## gin渲染  

### 各种数据格式的响应  

json、结构体、XML、YAML类似于java的properties、ProtoBuf  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "github.com/gin-gonic/gin/testdata/protoexample"
)

// 多种响应方式
func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // 1.json
   r.GET("/someJSON", func(c *gin.Context) {
      c.JSON(200, gin.H{"message": "someJSON", "status": 200})
   })
   // 2. 结构体响应
   r.GET("/someStruct", func(c *gin.Context) {
      var msg struct {
         Name    string
         Message string
         Number  int
      }
      msg.Name = "root"
      msg.Message = "message"
      msg.Number = 123
      c.JSON(200, msg)
   })
   // 3.XML
   r.GET("/someXML", func(c *gin.Context) {
      c.XML(200, gin.H{"message": "abc"})
   })
   // 4.YAML响应
   r.GET("/someYAML", func(c *gin.Context) {
      c.YAML(200, gin.H{"name": "zhangsan"})
   })
   // 5.protobuf格式,谷歌开发的高效存储读取的工具
   // 数组？切片？如果自己构建一个传输格式，应该是什么格式？
   r.GET("/someProtoBuf", func(c *gin.Context) {
      reps := []int64{int64(1), int64(2)}
      // 定义数据
      label := "label"
      // 传protobuf格式数据
      data := &protoexample.Test{
         Label: &label,
         Reps:  reps,
      }
      c.ProtoBuf(200, data)
   })

   r.Run(":8000")
}
```

### HTML模板渲染  

- gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据，本质上就是字符串替换  
- LoadHTMLGlob()方法可以加载模板文件  

### 重定向  

### 同步异步  

- goroutine机制可以方便地实现异步处理  
- 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "time"
   "log"
)

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // 1.异步
   r.GET("/long_async", func(c *gin.Context) {
      // 需要搞一个副本
      copyContext := c.Copy()
      // 异步处理
      go func() {
         time.Sleep(3 * time.Second)
         log.Println("异步执行：" + copyContext.Request.URL.Path)
      }()
   })
   // 2.同步
   r.GET("/long_sync", func(c *gin.Context) {
      time.Sleep(3 * time.Second)
      log.Println("同步执行：" + c.Request.URL.Path)
   })

   r.Run(":8000")
}
```

## gin中间件  

- gin可以构建中间件，但它只对注册过的路由函数起作用  
- 对于分组路由，嵌套使用中间件，可以限定中间件的作用范围  
- 中间件分为全局中间件，单个路由中间件和群组中间件  
- gin中间件必须是一个 gin.HandlerFunc 类型  

### 全局中间件  

- 所有请求都经过此中间件  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "time"
   "fmt"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {
   return func(c *gin.Context) {
      t := time.Now()
      fmt.Println("中间件开始执行了")
      // 设置变量到Context的key中，可以通过Get()取
      c.Set("request", "中间件")
      // 执行函数
      c.Next()
      // 中间件执行完后续的一些事情
      status := c.Writer.Status()
      fmt.Println("中间件执行完毕", status)
      t2 := time.Since(t)
      fmt.Println("time:", t2)
   }
}

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // 注册中间件
   r.Use(MiddleWare())
   // {}为了代码规范
   {
      r.GET("/middleware", func(c *gin.Context) {
         // 取值
         req, _ := c.Get("request")
         fmt.Println("request:", req)
         // 页面接收
         c.JSON(200, gin.H{"request": req})
      })

   }
   r.Run(":8000")
}
```

### Next()方法  

### 局部中间件  

### 中间件练习  

定义程序计时中间件，然后定义2个路由，执行函数后应该打印统计的执行时间，如下：  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "time"
   "fmt"
)

// 定义中间
func myTime(c *gin.Context) {
   start := time.Now()
   c.Next()
   // 统计时间
   since := time.Since(start)
   fmt.Println("程序用时：", since)
}

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // 注册中间件
   r.Use(myTime)
   // {}为了代码规范
   shoppingGroup := r.Group("/shopping")
   {
      shoppingGroup.GET("/index", shopIndexHandler)
      shoppingGroup.GET("/home", shopHomeHandler)
   }
   r.Run(":8000")
}

func shopIndexHandler(c *gin.Context) {
   time.Sleep(5 * time.Second)
}

func shopHomeHandler(c *gin.Context) {
   time.Sleep(3* time.Second)
}
```

## 会话控制  

### Cookie是什么  

- HTTP是无状态协议，服务器不能记录浏览器的访问状态，也就是说服务器不能区分两次请求是否由同一个客户端发出  
- Cookie就是解决HTTP协议无状态的方案之一，中文是小甜饼的意思  
- Cookie实际上就是服务器保存在浏览器上的一段信息。浏览器有了Cookie之后，每次向服务器发送请求时都会同时将该信息发送给服务器，服务器收到请求后，就可以根据该信息处理请求  
- Cookie由服务器创建，并发送给浏览器，最终由浏览器保存  

### Cookie的用途  

- 保持用户登录状态  
- 京东购物车  

### Cookie的使用  

测试服务端发送cookie给客户端，客户端请求时携带cookie  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "fmt"
)

func main() {
   // 1.创建路由
   // 默认使用了2个中间件Logger(), Recovery()
   r := gin.Default()
   // 服务端要给客户端cookie
   r.GET("cookie", func(c *gin.Context) {
      // 获取客户端是否携带cookie
      cookie, err := c.Cookie("key_cookie")
      if err != nil {
         cookie = "NotSet"
         // 给客户端设置cookie
         //  maxAge int, 单位为秒
         // path,cookie所在目录
         // domain string,域名
         //   secure 是否智能通过https访问
         // httpOnly bool  是否允许别人通过js获取自己的cookie
         c.SetCookie("key_cookie", "value_cookie", 60, "/",
            "localhost", false, true)
      }
      fmt.Printf("cookie的值是： %s\n", cookie)
   })
   r.Run(":8000")
}
```

### Cookie练习  

- 模拟实现权限验证中间件:有2个路由，login和home,login用于设置cookie,home是访问查看信息的请求,在请求home之前，先跑中间件代码，检验是否存在cookie
- 访问home，会显示错误，因为权限校验未通过  
- 然后访问登录的请求，登录并设置cookie  
- 再次访问home，访问成功  

```go
package main

import (
   "github.com/gin-gonic/gin"
   "net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
   return func(c *gin.Context) {
      // 获取客户端cookie并校验
      if cookie, err := c.Cookie("abc"); err == nil {
         if cookie == "123" {
            c.Next()
            return
         }
      }
      // 返回错误
      c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
      // 若验证不通过，不再调用后续的函数处理
      c.Abort()
      return
   }
}

func main() {
   // 1.创建路由
   r := gin.Default()
   r.GET("/login", func(c *gin.Context) {
      // 设置cookie
      c.SetCookie("abc", "123", 60, "/",
         "localhost", false, true)
      // 返回信息
      c.String(200, "Login success!")
   })
   r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
      c.JSON(200, gin.H{"data": "home"})
   })
   r.Run(":8000")
}
```

### Cookie的缺点  

- 不安全，明文  
- 增加带宽消耗  
- 可以被禁用  
- cookie有上限  

## Session中间件开发  

- 设计一个通用的Session服务，支持内存存储和redis存储  
- session模块设计:本质上k-v系统，通过key进行增删改查,session可以存储在内存或者redis（2个版本）  
- Session接口设计:Set(),Set(),Del(),Save()：session存储，redis的实现延迟加载  
- SessionMgr接口设计：Init()：初始化，加载redis地址,CreateSeesion()：创建一个新的session,GetSession()：通过sessionId获取对应的session对象  
- 
