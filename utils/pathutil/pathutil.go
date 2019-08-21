package pathutil

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// 获取当前文件执行的路径
func GetWorkingPath() string {
	file, _ := exec.LookPath(os.Args[0])

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	return rst
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 获取当前文件执行的路径
func GetGoPackagePath(pkg string) (goPath, pkgPath string, err error) {
	gopath := os.Getenv("GOPATH")
	arr := strings.Split(gopath, ";")
	for _, val := range arr {
		goPath = val
		pkgPath = val + "/src/" + pkg
		if PathExists(pkgPath) == true {
			return
		}
	}

	err = errors.New("not found")
	return
}
