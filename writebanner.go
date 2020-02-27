package main

import (
		"io/ioutil"
		"fmt"		
)

func Writebanner() []byte {
	// namefile := "Banner.txt"
	data, err := ioutil.ReadFile("Banner.txt")
	if err != nil {
		fmt.Println(" The name of file is not correct")
	}
	// fmt.Println(data)

	return data
}