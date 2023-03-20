package main

import (
	"log"
	"os"
	"time"
)

var (
	file *os.File
	err  error
)

func main() {
	/*
			file, err := os.Open("test2.txt")
			if err != nil {
				log.Fatal(err)
			}
			file.Close()
			// OpenFile提供更多的选项。
			// 最后一个参数是权限模式permission mode
			// 第二个是打开时的属性
			file, err = os.OpenFile("test2.txt", os.O_APPEND, 0666)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()
			fileInfo, err := os.Stat("test2.txt")
			if err != nil {
				if os.IsNotExist(err) {
					log.Fatal("File does not exist.")
				}
			}
			log.Println("File does exist. File information:")
			log.Println(fileInfo)
		// 检查读写权限
		file, err = os.OpenFile("test2.txt", os.O_WRONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				log.Println("Error: Write permission denied.")
			}
		}
		log.Println("Write permission.")
		file.Close()
		// 测试读权限
		file, err = os.OpenFile("test.txt", os.O_RDONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				log.Println("Error: Read permission denied.")
			}
		}
		log.Println("Read permission.")
		file.Close()
	*/
	err := os.Chmod("test2.txt", 0777)
	if err != nil {
		log.Println(err)
	}
	// 改变文件所有者
	err = os.Chown("test2.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}
	// 改变时间戳
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test2.txt", lastAccessTime, lastModifyTime)
	if err != nil {
		log.Println(err)
	}
}
