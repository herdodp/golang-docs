package main
import "fmt"

func main(){

	// type on go : int, float32, string, bool
	// two type of variable, var and const. var changeable, and const unchangeable
	// const variable can declare outside func main. 
	// example of const :

	const nametrue string = "nateguy"
	// nametrue string = "nategame" 
	// error

	//fmt.println(nametrue)
	// output is nateguy and unchangeable value



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
	var a,b,c string= "herdo ", "dimas ", "pratirto"

	fmt.Print(a, b, c)

	fmt.Print("\n")
	//in block
	var (
		angka1 int = 1
		angka2 int = 2
		angka3 int = 3
	)
	angka3 = 4 //change value of angka3 from 3 to 4

	fmt.Println(angka1, angka2, angka3)

	// write the output can using Print(), Println(), and printf()
	// Print() is default horizontal output, Println() is output with new line, and Printf() is given formating verbs
	// Printf() can print the formating verbs like %V for value, %T for data type, and others
	// we can using Print() with new line, example : Print("herdo\n") or Print("1", "\n")
	
}
