# code-any-test

通过下面代码，解析出`Hello`这个路由。

```go
package main

import (
	"github.com/gin-gonic/gin"
)

var prefixOfHello = "/test/hello"
var methodOfHello = "GET"

func handlerOfHello(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "hello",
	})
}
```