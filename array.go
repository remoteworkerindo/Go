package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "adri"
	names[1] = "meritus"
	names[2] = "tuhu"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		95,
		97,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(len(names))

}
