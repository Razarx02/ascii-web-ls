package main


import (
	"io/ioutil"
	"fmt"
)

// Main.
func Outputs( str string, namefile string) {
	text := GetAsArraystr(str)
	if namefile == ".txt" {
		namefile = "standart.txt"
	}
	toreturn := GetArray(text, namefile)
	PrintToBanner("Banner.txt", toreturn)
}

// Bytetostr.. used in main..
func Bytetostr(x []byte) string {
	txt := ""
	for i := range x {
		txt = txt + string(x[i])
	}
	return txt
}


// PrintToBanner prints the output into file (banner.txt)
func PrintToBanner(nameBanner string, ArrayOfoutput []byte) {

	toOutput := ioutil.WriteFile(nameBanner, ArrayOfoutput, 0644)
	if toOutput != nil {
		fmt.Println("Not correct")
	}
}

// GetAsArray function
func GetAsArraystr(str string) []string {
	Size := 0
	for range str {
		Size++
	}
	txt := ""
	Array := []string{}
	for i := 0; i <= Size-1; i++ {
		if str[i] == 13 && str[i+1] == 10 {
			Array = append(Array, txt)
			txt = ""
		} else if i != 0 && str[i-1] == 13 && str[i] == 10 {
			continue
		} else {
			txt = txt + string(str[i])
		}
	}
	Array = append(Array, txt)
	return Array
}


// Get Array of ascii art
func GetArray(ArgAsArray []string, namefil string) []byte  {
	ArrayOfoutput := []byte{}
	data, err := ioutil.ReadFile("./shablon/" + namefil)
	if err != nil {
		fmt.Println("ERROR")
	}
	for z := range ArgAsArray {
		for n :=-1; n < 7; n++ {
			for v := range ArgAsArray[z] {
				start := (int(ArgAsArray[z][v])-32)*9 + 2
				count := 0
				for a := range data {
					if data[a] == 10 {
						count++
					}
					if count == start+n {
						if data[a] != 10 {
								ArrayOfoutput = append(ArrayOfoutput, data[a])
							}
						}
					}

				}
				ArrayOfoutput = append(ArrayOfoutput, 10)

			}

		}
		return ArrayOfoutput
}
