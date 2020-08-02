package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	// 这里取绝对路径，方便打印出来的语法树可以转跳到编辑器
	f, err := parser.ParseFile(fset, "../demo.go", nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	// 打印语法树
	ast.Print(fset, f)
}
