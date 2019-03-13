package library

import (
	"io/ioutil"
)

//获取某目录下所有的文件, 不含子目录
func GetDirFiles(dir string) (list []string) {
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	for key, file := range info {
		if file.IsDir() == false {
			if dir[len(dir)-1:] != "/" {
				dir += "/"
			}
			list = append(list, dir+info[key].Name())
		}
	}
	return
}
