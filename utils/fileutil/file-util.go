package fileutil

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

var (
	DefaultIgnoreReg *regexp.Regexp = regexp.MustCompile(`/\.`)
)

func FormatPath(p string) string {
	p = strings.Replace(p, "\\", "/", -1)
	return p
}

func GetDirFiles(dirname string, ignoreDir []string, ignoreReg *regexp.Regexp) (files []string, err error) {
	var (
		dir  *os.File
		f    *os.File
		stat os.FileInfo
		list []os.FileInfo
	)
	//打开文件资源
	if dir, err = os.Open(dirname); err != nil {
		return
	}

	//如果dirname是一个文件，则直接返回该文件
	if stat, err = dir.Stat(); err != nil {
		return
	}

	if stat.IsDir() == false {
		err = errors.New("dirname必须是一个目录")
		return
	}

	//如果dirname是一个目录，则遍历目录
	stack := []string{dirname}
	dir.Close()
	for {
		if len(stack) == 0 {
			break
		}

		a := stack
		var x string
		x, a = a[0], a[1:]
		stack = a

		if f, err = os.Open(x); err != nil {
			return
		}

		if list, err = f.Readdir(0); err != nil {
			f.Close()
			return
		} else {
			for _, fi := range list {
				newName := x + "/" + fi.Name()
				if ignoreReg != nil {
					if ignoreReg.MatchString(fi.Name()) {
						continue
					}
					if ignoreReg.MatchString(newName) {
						continue
					}
				}

				if fi.IsDir() {
					if ignoreDir != nil {
						flag := false
						for _, val := range ignoreDir {
							ignore := strings.Index(newName, val)
							if ignore != -1 {
								flag = true
								break
							}
						}

						if flag == true {
							continue
						}
					}

					stack = append(stack, newName)
				} else {

					newName = FormatPath(newName)
					files = append(files, newName)
				}
			}
		}

		f.Close()
	}

	return
}
