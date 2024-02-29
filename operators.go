package main
import "fmt"

func main(){

// aritmathic using + - * / %

//common example
angka1 := 10
angka2 := 20
result := angka1 + angka2

fmt.Println("result first example is : ",result)

// in block
var(
	angka3 = 20
	angka4 = 20
)

result2 := angka3 + angka4
fmt.Println("result second example is : ",result2)

//comparison >,<,<=, >=
result_comparison := angka3<angka4
fmt.Println(result_comparison)

//logical &&, ||
result_logical := angka1<angka2 && angka3>angka4
fmt.Println(result_logical)




}