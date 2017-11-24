package brownfox

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/JodeZer/lazydog/file"
	"github.com/JodeZer/lazydog/inject"
)

type BrownFox struct {
	path   string
	dirs   []string
	deepth int
	inject *inject.Injector
	file.Jumper
}

func NewBrownFox(path string, deepth int) *BrownFox {
	return &BrownFox{
		path:   path,
		deepth: deepth,
		inject: inject.NewInjector(),
		dirs:   file.TreeDir(path, deepth),
	}
}

func (b *BrownFox) Backup() error {
	for _, dir := range b.dirs {
		b.BackupPath(dir)
	}
	return nil
}

func (b *BrownFox) Restore() error {
	for _, dir := range b.dirs {
		if err := b.RestorePath(dir); err != nil {
			return err
		}
		gofiles := file.ListGoFile(dir)

		parser := inject.NewParser(token.NewFileSet(), gofiles[0])
		if err := parser.Parse(); err != nil {
			return err
		}
		dgHelper := inject.NewDogHelper(dir, parser.PkgName())
		if err := dgHelper.EraseDogHelper(); err != nil && !os.IsNotExist(err) {
			return err
		}

	}

	return nil
}

func (b *BrownFox) Inject() error {
	for _, dir := range b.dirs {
		for _, gofile := range file.ListGoFile(dir) {

			fset := token.NewFileSet()
			parser := inject.NewParser(fset, gofile)
			err := parser.Parse()
			if err != nil {
				panic(err)
			}

			// inject
			parser.ForEachDecl(func(decl ast.Decl) {
				b.inject.InjectFunc(decl)
			})

			// write to new file
			var buf bytes.Buffer
			printer.Fprint(&buf, fset, parser.GetAst())
			//fmt.Println(buf.String())
			if err := ioutil.WriteFile(gofile, buf.Bytes(), os.ModeExclusive); err != nil {
				panic(err)
			}

			// write helper
			dgHelper := inject.NewDogHelper(dir, parser.PkgName())
			dgHelper.WriteDogHelper()
		}
	}

	return nil
}
