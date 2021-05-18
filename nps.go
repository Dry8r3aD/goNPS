package main

import (
	"container/list"
	"flag"
	"fmt"
	"os"
)

func checkSupportedTcType(TcType string) bool {
	supportedTypeList := list.New()
	supportedTypeList.PushBack("toml")
	supportedTypeList.PushBack("yaml")

	for i := supportedTypeList.Front(); i != nil; i = i.Next() {
		if i.Value == TcType {
			return true
		}
	}

	return false
}

func main() {
	// Command-Line arguments(Flag)
	tcFile := flag.String("tc", "", "Input testcase file's path")
	tcType := flag.String("type", "yaml", "Input testcase file's type")
	flag.Parse()

	if len(*tcFile) == 0 {
		fmt.Println("Invalid testcase filename")
		return
	}

	if checkSupportedTcType(*tcType) != true {
		fmt.Printf("Unsupported Testcase type (Type:%#v)\n", *tcType)
		return
	}

	if _, err := os.Stat(*tcFile); err != nil {
		fmt.Println("File does not Exist.")
		return
	}

	// 포인터 변수이므로 앞에 * 를 붙어 deference 해야
	fmt.Println(*tcFile)
}