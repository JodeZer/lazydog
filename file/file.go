package file

import (
	"io/ioutil"

	"github.com/jodezer/lazydog/inject"
)

type BrownFox struct {
	path   string
	deepth int
	inject *inject.Injector
	jumper
}

func NewBrownFox(path string, deepth int) *BrownFox {
	return &BrownFox{
		path:   path,
		deepth: deepth,
		inject: inject.NewInjector(),
	}
}

func treeDir(path string, deepth int) []string {
	paths := []string{path}
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, fi := range fis {
		if fi.IsDir() && deepth != 0 {
			nextDeepth := -1
			if deepth != -1 {
				nextDeepth = deepth - 1
			}
			paths = append(paths, treeDir(fi.Name(), nextDeepth)...)
		}
	}

	return paths
}

//===========
type jumper struct {
}

func (j *jumper) backupPath(paths string) error {
	return nil
}

func (j *jumper) restorePath(path string) error {
	return nil
}
