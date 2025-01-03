package file_vs_database

import (
	"fmt"
	"os"
)

/*
  在重命名之前将数据写入到磁盘上，但是元数据不一定会被顺利写入到磁盘
*/

func saveData3(path string, data []byte) error {

	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	// 0664 是文件的权限模式，表示文件所有者和所属组有读写权限，其他用户只有读权限。
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}

	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	// 在重命名之前将数据刷新到磁盘
	err = fp.Sync() // fsync
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)

}
