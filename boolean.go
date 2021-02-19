package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 85

	// fmt.Println(ujian)
	// fmt.Println(absensi)

	// var lulusUjian = ujian > 80
	// var lulusAbsensi = absensi > 80

	// var lulus = lulusUjian || lulusAbsensi
	fmt.Println(ujian >= 80 && absensi >= 80)

}
