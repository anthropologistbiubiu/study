package main

import "fmt"

type Inode interface {
	print(indentation string)
	clone() Inode
}

type file struct {
	name string
}

func (this *file) print(indentation string) {
	fmt.Println(indentation + this.name)
}
func (this *file) clone() Inode {
	return &file{name: this.name + "_clone"}
}

type folder struct {
	files []Inode
	name  string
}

func (this *folder) print(indentation string) {
	fmt.Println(indentation + this.name)
	for _, f := range this.files {
		f.print(indentation)
	}
}

func (this *folder) clone() Inode {
	var tem folder
	tem.name = this.name + "_clone"
	for _, f := range this.files {
		c := f.clone()
		tem.files = append(tem.files, c)
	}
	return &tem
}

func main() {

	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := folder{
		name:  "folder1",
		files: []Inode{file1, file2, file3},
	}
	file1_clone := file1.clone()
	file1_clone.print(" ")
	folder1_clone := folder1.clone()
	folder1_clone.print(" ")

}
