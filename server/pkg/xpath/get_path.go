// @author AlphaSnow

package xpath

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// GetPathByCaller
// 固定相对地址 /pkg/xpath/path.go
func GetPathByCaller() (p string, err error) {
	var ok bool
	if _, p, _, ok = runtime.Caller(0); ok == false {
		err = errors.New("not found runtime.Caller(0)")
		return
	}
	return filepath.Dir(filepath.Dir(filepath.Dir(p))), nil
}

// GetPathByExecutable
// 固定相对地址 ./main.go
func GetPathByExecutable() (p string, err error) {
	var ep string
	ep, err = os.Executable()
	if err != nil {
		return "", err
	}
	p = filepath.Dir(ep)
	return p, nil
}

// GetPathByWork
// 固定相对地址 ./
func GetPathByWork() (p string, err error) {
	return os.Getwd()
}
