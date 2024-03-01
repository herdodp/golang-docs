package main
import "fmt"


//declare struct
type person struct {
	nama string
	umur int
	profesi string
}



func main(){

var orang1 person
var orang2 person

//orang1 spesific
orang1.nama = "herdo"
orang1.umur = 30
orang1.profesi = "programmer"

//orang2 spesific
orang2.nama = "jhon doe"
orang2.umur = 25
orang2.profesi = "tukang somay"

//access data person orang1
fmt.Println("nama orang 1 : ", orang1.nama)

//access data person orang2
fmt.Println("nama orang 2 : ", orang2.nama)





}