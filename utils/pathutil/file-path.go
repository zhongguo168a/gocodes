package pathutil

import (
	"path/filepath"
	"strings"
)

type FilePath string

func (path FilePath) Name() (r string) {
	r = filepath.Base(string(path))
	ext := filepath.Ext(string(path))
	if ext != "" {
		r = strings.Replace(r, ext, "", -1)
	}
	return
}
