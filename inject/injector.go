package inject

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
)

const INJECT = `
package main

//import "fmt"

func a(){
	__traceStack()
}
`

type Injector struct {
	MyImport *ast.ImportSpec
	Mystate  ast.Stmt
}

func (i *Injector) InjectFunc(f ast.Decl) error {

	fd, ok := f.(*ast.FuncDecl)
	if !ok {
		return fmt.Errorf("not func")
	}
	newList := make([]ast.Stmt, 0, len(fd.Body.List)+1)

	newList = append(newList, i.Mystate)

	newList = append(newList, fd.Body.List...)

	fd.Body.List = newList
	return nil
}

func (i *Injector) InjectFile(path string) error {
	fset := token.NewFileSet()
	fbytes, err := ioutil.ReadFile(path)
	index := strings.LastIndex(path, `/`)
	f, err := parser.ParseFile(fset, string(path[index+1:]), fbytes, 0)
	if err != nil {
		return err
	}

	for _, decl := range f.Decls {
		i.InjectFunc(decl)
	}

	return nil
}

func NewInjector() *Injector {
	i := &Injector{}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", INJECT, 0)
	if err != nil {
		panic(err)
	}

	for _, d := range f.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok {
			i.Mystate = fd.Body.List[0]
		}
	}
	return i
}
