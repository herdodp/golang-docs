package main
import "fmt"

func main(){

	//type on go : int, float32, string, bool

	//type string
	var name1 string= "herdo"

	//type inferred
	var name2 = "dimas"

	//type inferred
	name3 := "pratirto"

	//inferred means compiler decide type based on the value

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(name3)



	//multiple variable 

	//one line
	var a,b,c string= "herdo", "dimas", "pratirto"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//in block
	var (
		angka1 int = 1
		angka2 int = 2
		angka3 int = 3
	)

	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(angka3)
	
}
