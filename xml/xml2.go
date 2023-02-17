package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

/*
	type People struct {
		XMLName xml.Name `xml:"people"`
		Id      int      `xml:"id,attr"`
		Name    string   `xml:"name"`
		Address string   `xml:"address"`
	}

	func main() {
		peo := &People{}
		b, err := ioutil.ReadFile("demo.xml")
		if err != nil {
			fmt.Println(err)
		}
		err = xml.Unmarshal(b, peo)
		fmt.Printf("%+v\n", peo)
		fmt.Println(peo.XMLName.Local)
		fmt.Println(peo.XMLName.Space)
		fmt.Println(peo.Id)
		fmt.Println(peo.Name)
		fmt.Println(peo.Address)
	}
*/
type Peoples struct {
	XMLName xml.Name `xml:"peoples"`
	Version string   `xml:"version,attr"`
	Peos    []People `xml:"people"`
}

type People struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {

	data, err := ioutil.ReadFile("demo1.xml")
	if err != nil {
		fmt.Println(err)
	}
	peo := &Peoples{}
	err = xml.Unmarshal(data, peo)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", peo)
}
