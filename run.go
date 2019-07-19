package main

import (
	"fmt"
	"github.com/robteix/testmod"
	"io/ioutil"
	"strings"
	"toolkit/demo"
	"toolkit/druid"
	"toolkit/httpserver"
	"toolkit/json"
)

func main() {
	druid.GetDspFeeStat()
	//fmt.Println("GitUrl: ", demo.GetExistedGitUrl("/e/Projects/Go/toolkit"))
	//testJson()

	httpserver.StartServer()

	fmt.Println(testmod.Hi("roberto"))

	lines := loadFile("C:\\Users\\wxp04\\Downloads\\emi_ad_app_0301-1.txt")
	fmt.Println("lines number: ", len(lines))
	for _,line := range lines {
		adid := line[0:9]
		//fmt.Println(line[10:])
		m, n := demo.ParseProduct(line[10:])
		fmt.Printf("%s,%d,%s\n",adid, m, n)
	}
}

func loadFile(file string) []string {
	dat, _ := ioutil.ReadFile(file)
	var str = string(dat)
	//fmt.Println(str)
	return strings.Split(str, "\n")
}

func testJson() {
	json.DecodeJson()
}