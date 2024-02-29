package main
import "fmt"

func main(){

	//loops

	//for loops
	//example 1. iterasi 0 sampai 9
	for i:=0; i<10; i++{
		fmt.Println(i)
	}

	//example 2. print bilangan genap
	for i:=0; i<10; i+=2{
		fmt.Println("GENAP KE : ",i)
	}

	//example 3. skip number 5
	for x:=0; x<10; x++{
		if(x == 5){
			continue
		}

		fmt.Println("angka ke ",x)

	}

	//example 4. stop the iteration when touch number 5
	for x:=0; x<10; x++{
		if(x == 5){
			break
		}

		fmt.Println("angka : ",x)

	}







}