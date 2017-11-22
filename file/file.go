package file

import (
	"io/ioutil"
	"strings"

	"github.com/jodezer/lazydog/inject"
)

const BackupSuffix = ".ld"

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

func Backup(b *BrownFox) error {
	return nil
}

func Restore(b *BrownFox) error {
	return nil
}

func Inject(b *BrownFox) error {
	return nil
}

//===== funcs

func hiddenDir(path string) bool { // just tmp impl
	sps := strings.Split(path, `/`)
	if len(sps) == 0 {
		return true
	}
	last := sps[len(sps)-1]

	if len(last) == 0 {
		return true
	}

	if len(last) == 1 {
		return false
	}

	return last[0] == '.' && last[1] != '.'
}

func listGoFile(path string) []string {
	// fis, err := ioutil.ReadDir(path)
	// if err != nil {
	// 	panic(err)
	// }
	// farray := []string{}
	// for _, fi := range fis {
	// 	if fi.IsDir() {
	// 		continue
	// 	}
	// 	if strings.HasSuffix(fi.Name(), ".go") && !strings.HasSuffix(fi.Name(), "_test.go") {
	// 		farray = append(farray, path+"/"+fi.Name())
	// 	}
	// }
	// return farray
	return listSuffixFile(path, []string{".go"}, "_test.go")
}

func listSuffixFile(path string, include []string, exclude ...string) []string {
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	farray := []string{}
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		for _, exclu := range exclude {
			if strings.HasSuffix(fi.Name(), exclu) {
				continue
			}
		}
		for _, inclu := range include {
			if strings.HasSuffix(fi.Name(), inclu) {
				farray = append(farray, path+"/"+fi.Name())
			}
		}
	}
	return farray
}

func listGofileByPaths(paths []string) []string {
	ret := []string{}
	for _, path := range paths {
		fs := listGoFile(path)
		ret = append(ret, fs...)
	}
	return ret
}

func treeDir(path string, deepth int) []string {
	if hiddenDir(path) {
		return []string{}
	}
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
			paths = append(paths, treeDir(path+"/"+fi.Name(), nextDeepth)...)
		}
	}

	return paths
}

func backupFileName(fileName string) string {
	sps := strings.Split(fileName, `/`)
	backupFileName := sps[len(sps)-1] + BackupSuffix
	sps[len(sps)-1] = backupFileName
	return strings.Join(sps, "/")
}

func restoreFileName(fileName string) string {
	sps := strings.Split(fileName, `/`)
	if !strings.HasSuffix(fileName, BackupSuffix) {
		return ""
	}
	restoreFileName := strings.Replace(sps[len(sps)-1], BackupSuffix, "", 1)
	sps[len(sps)-1] = restoreFileName
	return strings.Join(sps, "/")
}

func copyFile(src string, dst string) error {

	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, data, 0644)
}

//===========
type jumper struct {
}

func (j *jumper) backupPath(path string) error {
	files := listGoFile(path)
	for _, fn := range files {
		if err := copyFile(fn, backupFileName(fn)); err != nil {
			return err
		}
	}
	return nil
}

func (j *jumper) restorePath(path string) error {
	files := listSuffixFile(path, []string{".go" + BackupSuffix})
	for _, fn := range files {
		if err := copyFile(fn, restoreFileName(fn)); err != nil {
			return err
		}
	}
	return nil
}
