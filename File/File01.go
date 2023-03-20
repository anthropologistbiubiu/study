package main

import (
	"fmt"
	"log"
	"os"
)

var (
	file     *os.File
	fileInfo os.FileInfo
	err      error
)

// 创建文件
func main() {
	file, err = os.Create("tes.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.Println(file)
	err = os.Truncate("test.txt", 100)
	if err != nil {
		log.Println(err)
	}
	// 裁剪一个文件到100个字节。
	// 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	// 如果文件本来超过100个字节，则超过的字节会被抛弃。
	// 这样我们总是得到精确的100个字节的文件。
	// 传入0则会清空文件。
	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

	originalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
	/*
		err = os.Remove("test2.txt")
		if err != nil {
			log.Fatal(err)
		}
	*/

}
