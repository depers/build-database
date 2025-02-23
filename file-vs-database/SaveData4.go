package file_vs_database

import (
	"os"
)

/**
仅附加日志的优点在于它不会修改现有数据，也不会处理重命名操作
这里又会诞生新的问题
1. 仅靠记录日志，我们如何解决数据查询的问题，日志什么时候刷新到磁盘上去。
2. 日志如何删除已有的日志文件，日志不能永远的增长下去。
*/

func LogCreate(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0664)
}

func LogAppend(fp *os.File, line string) error {
	buf := []byte(line)
	buf = append(buf, '\n')
	_, err := fp.Write(buf)
	if err != nil {
		return err
	}

	return fp.Sync()
}
