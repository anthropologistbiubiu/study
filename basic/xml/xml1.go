package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type XMLServers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Servers []Server `xml:"server"`
	//Description string   `xml:",innerxml"`
}

func main() {
	file, err := os.Open("./server.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	v := XMLServers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println()
		return
	}
	fmt.Printf("%+v\n", v)
}
