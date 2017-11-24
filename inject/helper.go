package inject

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	pkgPath "path"
	"runtime"
	"strings"

	"github.com/JodeZer/lazydog/file"
)

type DogHelper struct {
	path string
	pkg  string
}

func NewDogHelper(path, pkg string) *DogHelper {
	return &DogHelper{
		path: path,
		pkg:  pkg,
	}
}
func (d *DogHelper) WriteDogHelper() {

	_, filename, _, _ := runtime.Caller(0)

	fset := token.NewFileSet()
	fbytes, err := ioutil.ReadFile(pkgPath.Join(pkgPath.Dir(filename)) + "/dogHelper.go")
	if err != nil {
		panic(err)
	}
	f, err := parser.ParseFile(fset, "dogHelper.go", fbytes, 0)
	if err != nil {
		panic(err)
	}

	f.Name.Name = d.pkg

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, f)
	if err := ioutil.WriteFile(genPath(d.path, d.pkg), buf.Bytes(), os.ModePerm); err != nil {
		panic(err)
	}

}

func (d *DogHelper) EraseDogHelper() error {
	return os.Remove(genPath(d.path, d.pkg))
}

func genPath(path, pkg string) string {
	suffix := ""
	if !strings.HasSuffix(path, "/") && len(path) != 0 {
		suffix = "/"
	}
	return path + suffix + "gen_" + pkg + file.HelperSuffix
}
