package file_vs_database

import (
	"os"
)

/**
仅附加日志的优点在于它不会修改现有数据，也不会处理重命名操作
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
