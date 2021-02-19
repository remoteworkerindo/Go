package main

import "fmt"

func main() {
	var name = "jaswan"

	if name == "adri" {
		fmt.Println("Hello Adri")
	} else if name == "meritus" {
		fmt.Println("Hello Meritus")
	} else if name == "tuhu" {
		fmt.Println("Hello Tuhu")
	} else {
		fmt.Println("kenalan dong!!")
	}

	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

}
