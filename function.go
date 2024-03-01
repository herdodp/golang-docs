package main
import "fmt"

// func method without paramater
func nama(){
	fmt.Println("herdo dimas pratirto")
}

//example function with parameter
func namaumur(nama string, umur int){
	fmt.Println("nama saya ",nama, " dan umur saya ", umur)
}



//example return values
func penjumlahan(angka1 int, angka2 int) int {
	return angka1 + angka2
}


//example named return values 
func perkalian(angka1 int, angka2 int)(hasil int){
	hasil = angka1 * angka2
	return
}

//example for multiple return values
func kalijumlah(angka1 int, angka2 int)(hasil int, result string){
 hasil = angka1 + angka2
 result = "ini adalah hasil "
 return
}


//example for recursion
func testcount(x int) int {
	if x == 11 {
	  return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
  }


//example 2 for recursion
func count(i int)int{
	if i == 10{
		return 0
	}
	fmt.Println(i)
	return count(i + 1)
}






func main(){

	/*
	example multiple call function
	for x:=0; x<10; x++{
		nama()
	}
	*/

	/*
	example for function with parameter
	namaumur("herdo", 30)
	*/

	/*
	example for return values
	fmt.Println(penjumlahan(10,10))
	*/

	
	/*example for name return values
	total := perkalian(20,20)
	fmt.Println(total)
	*/
	
	/*
	example for multiple return
	cetak, hasil := kalijumlah(10,10)
	fmt.Println(cetak, hasil)
	*/

	/*
	example 1 for recursion 
	testcount(1)
	*/

	/*
	example 2 for recursion
	count(1)
	*/

	

}