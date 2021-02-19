package main

import "fmt"

func main() {
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("perulangan ke", counter)
	// }

	slice := []string{"adri", "meritus", "tuhu", "jaswan", "andri"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	for i, nilai := range slice {
		fmt.Println("index", i, "=", nilai)
	}

	for _, nilai := range slice {
		fmt.Println(nilai)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
