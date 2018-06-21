package main

import (
	"fmt"
	"io/ioutil"
)

func main(){
	datFile1,errFile1 := ioutil.ReadFile("fileRead.txt")
	if errFile1 != nil {
		panic(errFile1)
	}

	datFile2, errFile2 := ioutil.ReadFile("questions.txt")
	if errFile2 != nil {
		panic(errFile2)
	}

	fmt.Print(string(datFile1),string(datFile2))
}