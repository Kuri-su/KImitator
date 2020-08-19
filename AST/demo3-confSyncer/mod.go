package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	GofileParseDemo()
	//GomodParseDemo()

}

func GofileParseDemo() {
	var filename = "/home/kurisucode/myRepo/KImitator-and-studies/AST/demo3-confSyncer/confSyncer"
	files, err := GetAllFiles(filename)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		f, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.AllErrors)
		if err != nil {
			panic(err)
		}

		fmt.Println(f)
		fmt.Println(file)
	}

}

// 去获取 项目层面 的 依赖
func GomodParseDemo() {
	var filename = "/home/kurisucode/myRepo/KImitator-and-studies/AST/demo3-confSyncer/confSyncer/go.mod"

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lax, err := modfile.ParseLax("file", bytes, nil)
	if err != nil {
		panic(err)
	}

	for _, require := range lax.Require {
		fmt.Println(require.Mod.Path, require.Mod.Version)
	}
}
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}
