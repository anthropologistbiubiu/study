package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address   Address
		Comment   string `xml:",comment"`
	}
	v := &Person{
		//XMLName: "people",
		Id:        1,
		FirstName: "sunweiming",
		LastName:  "yujinling",
		Age:       18,
		Height:    180,
		Married:   false,
		Address: Address{
			City:  "香港",
			State: "single",
		},
		Comment: "摇滚万岁",
	}
	output, err := xml.MarshalIndent(v, "", "	")
	//output, err := xml.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(output)
	/*
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		os.Stdout.Write(output)
	*/
}
