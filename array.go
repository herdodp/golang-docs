package main
import "fmt"

func main(){

	// array inferred length
	var array1  = [...]string{"herdo", "dimas", "pratirto"}
	fmt.Println(array1[1])

	// array 4 length
	angka1 := [4]int{1,2,3,4}
	fmt.Println(angka1[2])

	// change value
	angka1[3] = 5
	fmt.Println(angka1)

	//not initial, partially initial, fully initial, specific initial
	var angka2 = [4]int{} // not initial
	var angka3 = [4]int{1,2,3} // partially initial
	var angka4 = [4]int{1,2,3,4} // fully initial
	var angka5 = [4]int{0:1,3:4} // specific intial. in element using index:value . in ex, 0:1. 0 is index, 1 is value.


	fmt.Println(angka2)
	fmt.Println(angka3)
	fmt.Println(angka4)
	fmt.Println(angka5)

	//to know lenght of array
	fmt.Println("length array1 : ", len(array1))

	//slice array
	var array6 = []int{1,2,3,4,5,6,7}
	fmt.Println("slice array : ", array6[1:4])


	//append element inside array
	array6 = append(array6,8,9,10)
	fmt.Println(array6)


	//combine array
	var array7 = []int{1,2,3}
	var array8 = []int{4,5,6}
	var array9 = []int{7,8,9}

	var array10 = append(append(array7,array8...),array9...) //combine array 3 or more variable
	var array11 = append(array7,array8...) //combine array 2 variable

	fmt.Println(array10)
	fmt.Println(array11)


	//change value of array
	var array12 = []int{1,2,3,4,5} //original array
	array13 := array12[0:3] //slice length array

	array13 = append(array13,6,7) // change value of array12 from 4,5 to 6,7

	fmt.Println(array13)



}