package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	wirtenNum int
)

func main() {
	// 打开原始文件
	originalFile, err := os.Open("test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFle.Close()
	// 从源中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)
	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile(
		"test2.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// 写字节到文件中
	byteSlice := []byte("Bytes!\n")
	wirtenNum, err = file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", wirtenNum)

	err = ioutil.WriteFile("test3.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}

	//ioutil包有一个非常有用的方法WriteFile()可以处理创建／打开文件、写字节slice和关闭文件一系列的操作。
	//如果你需要简洁快速地写字节slice到文件中，你可以使用它。

}
