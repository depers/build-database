package file_vs_database

import "os"

/*
	 持久化数据到文件方法一：直接将数据写入文件中
	 缺点：
		1.直接清空原来的数据，使得原来的数据无法正常被并发读取
		2.写入文件的操作不是原子性的，并发读导致数据不完整
		3.无法保证数据会被持久化到磁盘上，数据仍有可能存储在操作系统的页面缓存中
*/
func SaveData1(path string, data []byte) error {

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)
	return err
}
