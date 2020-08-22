package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go/parser"
	"go/token"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"io/ioutil"
	"os"
	"strings"
)

var Result = struct {
	Imports map[string]bool
	ModDeps []module.Version
}{
	Imports: make(map[string]bool),
}

var (
	pkgDir   = "/home/kurisucode/myRepo/KImitator-and-studies/AST/demo3-confSyncer/confSyncer"
	pkgGoMod = "/home/kurisucode/myRepo/KImitator-and-studies/AST/demo3-confSyncer/confSyncer/go.mod"
)

func main() {
	GomodParseDemo()
	GofileParseDemo()
}

func GofileParseDemo() {
	var filename = pkgDir
	files, err := GetAllFiles(filename)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		f, err := parser.ParseFile(token.NewFileSet(), file, nil, parser.AllErrors)
		if err != nil {
			panic(err)
		}

		for _, impo := range f.Imports {
			importPkg := impo.Path.Value
			importPkg = string(importPkg[1 : len(importPkg)-1])

			if !CheckImportIsHimSelf(importPkg) {
				continue
			}

			if _, ok := packages[importPkg]; ok {
				continue
			}

			if _, ok := Result.Imports[importPkg]; !ok {
				Result.Imports[importPkg] = true
			}

		}

	}

	marshal, err := jsoniter.Marshal(Result)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))

}

// 去获取 项目层面 的 依赖
func GomodParseDemo() {
	var filename = pkgGoMod

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lax, err := modfile.ParseLax("file", bytes, nil)
	if err != nil {
		panic(err)
	}

	GoModPackageName = lax.Module.Mod.Path

	for _, require := range lax.Require {
		Result.ModDeps = append(Result.ModDeps, require.Mod)
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
