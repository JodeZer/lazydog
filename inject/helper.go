package inject

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
)

// type DogHelper struct {
// }

func WriteDogHelper(path, pkg string) {
	fset := token.NewFileSet()
	fbytes, err := ioutil.ReadFile("inject/dogHelper.go")
	if err != nil {
		panic(err)
	}
	f, err := parser.ParseFile(fset, "dogHelper.go", fbytes, 0)
	if err != nil {
		panic(err)
	}

	f.Name.Name = pkg

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, f)
	if err := ioutil.WriteFile(genPath(path, pkg), buf.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}

}

func genPath(path, pkg string) string {
	suffix := ""
	if !strings.HasSuffix(path, "/") && len(path) != 0 {
		suffix = "/"
	}
	return path + suffix + "gen_" + pkg + "_dgh.go"
}
