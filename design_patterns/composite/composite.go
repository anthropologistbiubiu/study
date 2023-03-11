package main

import "fmt"

type Component interface {
	search(key string) bool
}

type File struct {
	fileName string
	keyword  string
}

func (this *File) search(keyword string) bool {
	fmt.Printf("start search keyword %s in file %s\n", keyword, this.fileName)
	if this.keyword == keyword {
		fmt.Printf(" keyword %s in file %s\n", keyword, this.fileName)
		return true
	}
	return false
}

func (this *File) getname() string {
	return this.fileName
}

type Folder struct {
	Components []Component
	folderName string
}

func (this *Folder) search(keyword string) bool {
	fmt.Printf("start search keyword %s in Folder %s\n", keyword, this.folderName)
	for _, c := range this.Components {
		if c.search(keyword) {
			fmt.Println("search success")
			return true
		} else {
			continue
		}
	}
	fmt.Println("search failed")
	return false
}
func (this *Folder) AddFile(c Component) {
	this.Components = append(this.Components, c)
}

func main() {
	f1 := &File{
		fileName: "f1",
		keyword:  "sun",
	}
	f2 := &File{
		fileName: "f2",
		keyword:  "wei",
	}

	f3 := &File{
		fileName: "f3",
		keyword:  "ming",
	}
	folder := &Folder{
		folderName: "folder1",
	}
	folder.AddFile(f1)
	folder.AddFile(f2)
	folder.AddFile(f3)
	folder.search("swei")

}
