package file_vs_database

import (
	"fmt"
	"math/rand"
	"os"
)

// SaveData2 采用重命名的方式更新文件
func SaveData2(path string, data []byte) error {

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

	// 首先重命名操作是原子性的
	// 其次程序无法控制数据何时持久化到磁盘上，如果在重命名时程序崩溃，此时有可能文件的元信息（文件大小）可能已经被持久化到了磁盘上，但是数据却没有持久化到磁盘，此时数据文件就被损坏了
	return os.Rename(tmp, path)
}

func randomInt() int {

	return rand.Int()
}
