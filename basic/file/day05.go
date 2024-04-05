package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// 读取最多N 个字节

var (
	file *os.File
	err  error
)

func main1() {

	file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println("Open File err: ", err)
	}

	bytesRead := make([]byte, 10)

	numByteRead, err := file.Read(bytesRead)

	if err != nil {
		fmt.Println("Read err", err)
	}
	fmt.Printf("%s  %d\n", bytesRead, numByteRead)

}

// 读取正好N个字节
func main2() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	// file.Read()可以读取一个小文件到大的byte slice中，
	// 但是io.ReadFull()在文件的字节数小于byte slice字节数的时候会返回错误
	byteSlice := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of bytes read: %d\n", numBytesRead)
	fmt.Printf("Data read: %s\n", byteSlice)
}

// 读取最少N个字节
func main3() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	byteSlice := make([]byte, 512)
	minBytes := 20
	// io.ReadAtLeast()在不能得到最小的字节的时候会返回错误，但会把已读的文件保留
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes read: %d\n", numBytesRead)
	fmt.Printf("Data read: %s\n", byteSlice)
}

// 读取全部字节
func main4() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	// os.File.Read(), io.ReadFull() 和
	// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s", data)
	fmt.Println("Number of bytes read:", len(data))
}

// 快速读取
func main5() {
	// 读取文件到byte slice中
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read: %s", data)
}

// 使用缓存读取

func main() {
	// 打开文件，创建buffered reader
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	bufferedReader := bufio.NewReader(file)
	// 得到字节，当前指针不变
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)
	// 读取，指针同时移动
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)
	// 读取一个字节, 如果读取不成功会返回Error
	myByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read 1 byte: %c\n", myByte)
	// 读取到分隔符，包含分隔符，返回byte slice
	/*
		dataBytes, err := bufferedReader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Read bytes: %s\n", dataBytes)
	*/
	// 读取到分隔符，包含分隔符，返回字符串
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read string: %s\n", dataString)
	//这个例子读取了很多行，所以test.txt应该包含多行文本才不至于出错
}
