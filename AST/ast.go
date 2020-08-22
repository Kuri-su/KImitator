package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	// 这里取绝对路径，方便打印出来的语法树可以转跳到编辑器
	f, err := parser.ParseFile(fset, "./confSyncer/cmd/confsyncer/main.go", nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(f)
}

type Visitor struct {
}

func (v Visitor) Visit(node ast.Node) (w ast.Visitor) {
	return v
}

// print all actual dependent libs
func PrintAllActualDependentLibs(f *ast.File, fSet *token.FileSet) {
	v := &Visitor{}
	ast.Walk(v, f)
}
