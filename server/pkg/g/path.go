package g

import (
	"path"
	"server/pkg/xpath"
)

func Path(ps ...string) string {
	var rp string
	var err error
	if globalContainer.Has("path") {
		rp = globalContainer.Get("path").(string)
	} else {
		// 判断目录下有没指定文件判断地址是否正确
		rp, err = xpath.GetRootPathByLayoutFolder("config.yaml")
		if err != nil {
			panic(err)
		}
		globalContainer.Set("path", rp)
	}

	if len(ps) == 0 {
		return rp
	}

	nps := make([]string, len(ps)+1)
	nps[0] = rp
	copy(nps[1:], ps)
	return path.Join(nps...)
}
