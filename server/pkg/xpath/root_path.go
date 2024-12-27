// @author AlphaSnow

package xpath

import (
	"os"
	"path"
)

func GetRootPathByLayoutFolder(folders ...string) (string, error) {
	if rp, err := GetPathByExecutable(); err != nil {
		return "", err
	} else {
		if allRootFolderExists(rp, folders...) {
			return rp, nil
		}
	}

	return GetPathByCaller()
}

func allRootFolderExists(rootPath string, folders ...string) bool {
	for _, folder := range folders {
		if folderExists(path.Join(rootPath, folder)) == false {
			return false
		}
	}
	return true
}

func folderExists(folderPath string) bool {
	_, err := os.Stat(folderPath)
	if err == nil {
		return true
	}
	return false
}
