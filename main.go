package main

import (
	"go/parser"
	"go/token"
	"strings"
)

func main() {
	// src is the input for which we want to print the AST.
	src := `
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
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)

	if err != nil {
		panic(err)
	}

	collection := routerItemMap{}
	traverseCallback := func(name, t string) {
		println(name, t)
		if strings.HasPrefix(name, "prefixOf") {
			collection.collect(name[8:], "prefix")
		} else if strings.HasPrefix(name, "methodOf") {
			collection.collect(name[8:], "method")
		} else if strings.HasPrefix(name, "handlerOf") {
			collection.collect(name[9:], "router")
		}
	}
	traverse(f, traverseCallback)

	dumpCallback := func(name string) {
		println("to dump it to template => ", name)
	}
	collection.dump(dumpCallback)
	// var buf bytes.Buffer
	// printer.Fprint(&buf, fset, f)
	// println(buf.String())
}

// prefixOfXXX
// methodOfXXX
// handlerOfXXX
