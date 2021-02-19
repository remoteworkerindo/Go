import "fmt"

func main() {
	var name = "jaswan"

	// if name == "adri" {
	// 	fmt.Println("Hello Adri")
	// } else if name == "meritus" {
	// 	fmt.Println("Hello Meritus")
	// } else if name == "tuhu" {
	// 	fmt.Println("Hello Tuhu")
	// } else {
	// 	fmt.Println("kenalan dong!!")
	// }

	// if length := len(name); length > 5 {
	// 	fmt.Println("nama terlalu panjang")
	// } else {
	// 	fmt.Println("nama sudah benar")
	// }

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false
		fmt.Println("nama sudah benar")
	}


}
