package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"strconv"
)

// Visitor
type Visitor struct {
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	switch node.(type) {
	case *ast.GenDecl:
		genDecl := node.(*ast.GenDecl)
		// 查找有没有 import context 包
		// Notice：没有考虑没有 import 任何包的情况
		if genDecl.Tok == token.IMPORT {
			v.addImport(genDecl)
			// 不需要再遍历子树
			return nil
		}
	case *ast.InterfaceType:
		// 遍历所有的接口类型
		iface := node.(*ast.InterfaceType)
		addContext(iface)
		// 不需要再遍历子树
		return nil
	}
	return v
}

// addImport 引入 context 包
func (v *Visitor) addImport(genDecl *ast.GenDecl) {
	// 是否已经 import
	hasImported := false
	for _, v := range genDecl.Specs {
		imptSpec := v.(*ast.ImportSpec)
		// 如果已经包含"context"
		if imptSpec.Path.Value == strconv.Quote("context") {
			hasImported = true
		}
	}
	// 如果没有 import context，则 import
	if !hasImported {
		genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: strconv.Quote("context"),
			},
		})
	}
}

// addContext 添加 context 参数
func addContext(iface *ast.InterfaceType) {
	// 接口方法不为空时，遍历接口方法
	if iface.Methods != nil || iface.Methods.List != nil {
		for _, v := range iface.Methods.List {
			ft := v.Type.(*ast.FuncType)
			hasContext := false
			// 判断参数中是否包含 context.Context 类型
			for _, v := range ft.Params.List {
				if expr, ok := v.Type.(*ast.SelectorExpr); ok {
					if ident, ok := expr.X.(*ast.Ident); ok {
						if ident.Name == "context" {
							hasContext = true
						}
					}
				}
			}
			// 为没有 context 参数的方法添加 context 参数
			if !hasContext {
				ctxField := &ast.Field{
					Names: []*ast.Ident{
						ast.NewIdent("ctx"),
					},
					// Notice: 没有考虑 import 别名的情况
					Type: &ast.SelectorExpr{
						X:   ast.NewIdent("context"),
						Sel: ast.NewIdent("Context"),
					},
				}
				list := []*ast.Field{
					ctxField,
				}
				ft.Params.List = append(list, ft.Params.List...)
			}
		}
	}
}

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "./demo2.go", nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}

	// 遍历节点
	v := &Visitor{}
	ast.Walk(v, f)

	var output []byte
	buffer := bytes.NewBuffer(output)
	err = format.Node(buffer, fset, f)
	if err != nil {
		log.Fatal(err)
	}
	// 输出 Go 代码
	fmt.Println(buffer.String())

}
