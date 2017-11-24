package main

import (
	"github.com/JodeZer/lazydog/brownfox"
)

func main() {

	// inject.WriteDogHelper("example", "main")

	// fset := token.NewFileSet()
	// fbytes, err := ioutil.ReadFile("example/example_main.go")
	// f, err := parser.ParseFile(fset, "example_main.go", fbytes, 0)
	// if err != nil {
	// 	log.Fatal(err) // parse error
	// }

	// inj := inject.NewInjector()
	// for _, d := range f.Decls {

	// 	if fd, ok := d.(*ast.FuncDecl); ok {

	// 		inj.Inject(fd)

	// 	}
	// }

	// var buf bytes.Buffer
	// printer.Fprint(&buf, fset, f)
	// fmt.Println(buf.String())
	// if err := ioutil.WriteFile("example/example_main.go", buf.Bytes(), os.ModeExclusive); err != nil {
	// 	panic(err)
	// }

	bf := brownfox.NewBrownFox("example", -1)
	// if err := bf.Backup(); err != nil {
	// 	panic(err)
	// }
	// if err := bf.Inject(); err != nil {
	// 	panic(err)
	// }
	if err := bf.Restore(); err != nil {
		panic(err)
	}
}
