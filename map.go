package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Adri",
		"address": "Sulawesi",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := map[string]string{
		"title":  "Belajar Go-Lang",
		"author": "Tuhu",
		"ups":    "Wrong Books",
	}
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
