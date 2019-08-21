package pathutil

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// go的包
// 不包含gopath的包字符串标识
type GoPkg string

func (pkg GoPkg) MustGoPath() string {
	p, _ := pkg.GetGoPath()

	return p
}

// 获取go包所在的go path
func (pkg GoPkg) GetGoPath() (goPath string, err error) {
	envPath := os.Getenv("GOPATH")
	if envPath == "" {
		err = errors.New("get env $GOPATH: not found")
		return
	}

	arr := strings.Split(envPath, ";")
	for _, val := range arr {
		val = filepath.ToSlash(val)
		pkgPath := val + "/src/" + string(pkg)
		if PathExists(pkgPath) == false {
			continue
		}

		goPath = val
		return
	}

	err = errors.New("not found")
	return
}

func (pkg GoPkg) Name() string {
	return FilePath(pkg).Name()
}

// 获取完整的pkg文件路径
func (pkg GoPkg) GetFilePath() string {
	gopath, _ := pkg.GetGoPath()
	return gopath + "/src/" + string(pkg)
}

// 获取指定项目所在包的子包
func (pkg GoPkg) GetSubpackages() (subs []string, err error) {
	gopath, err := pkg.GetGoPath()
	if err != nil {
		return
	}

	submap := map[string]bool{}
	err = filepath.Walk(pkg.GetFilePath(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() == false {
			return nil
		}

		path = filepath.ToSlash(path)
		if strings.Contains(path, "generate") {
			return nil
		}
		if strings.Contains(path, "internal") {
			return nil
		}

		match, err := regexp.MatchString(`\/\.`, path)
		if match {
			return nil
		}
		pkg := strings.Replace(path, gopath+"/src/", "", -1)
		submap[pkg] = true
		return nil
	})
	if err != nil {
		return
	}

	subs = func() (x []string) {
		for key, _ := range submap {
			x = append(x, key)
		}

		sort.Strings(x)
		return
	}()

	return
}
