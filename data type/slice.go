package main

import "fmt"

func main() {
	names := [...]string{"Rian", "Ari", "Mulzahrian", "One", "own", "owl"}
	slice := names[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	names[5] = "has change"
	fmt.Println(slice)

	slice[0] = "Za"
	fmt.Println(names)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // [Senin, Selasa, Rabu ,Kamis, Jumat, Sabtu, Minggu Baru]

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) //[Ups, Minggu Baru, Libur Baru]
	fmt.Println(days)      // [Senin, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru]

	newSlice := make([]string, 2, 5)
	newSlice[0] = "One"
	newSlice[1] = "Art"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)
}
