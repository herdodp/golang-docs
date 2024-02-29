package main
import "fmt"

func main(){
	
	angka1 := 10
	angka2 := 20

	// if method

	if (angka1<angka2) {
		fmt.Println(angka1," lebih kecil ", angka2)
	}else if(angka1>angka2){
		fmt.Println(angka1, " lebih besar ", angka2)
	}else{
		fmt.Println(angka1, " sama dengan ", angka2)
	}


	//switch method
	//single case
	days := "wednesday"

	switch days {
	case "monday" : 
		fmt.Println("monday")

	case "tuesday" :
		fmt.Println("tuesday")
	
	case "wednesday" :
		fmt.Println("wednesday")
	
	case "thursday" :
		fmt.Println("thursday")
	
	case "friday" :
		fmt.Println("friday")
	
	case "saturday" :
		fmt.Println("saturday")
	
	case "sunday" :
		fmt.Println("sunday")
	
	default :
		fmt.Println("not a day")
	
	}


	// multi case
	number := 1
	
	switch number{
	case 1,2,3 :
		fmt.Println("positif number")
	case -1,-2,-3 :
		fmt.Println("negatif number")
	default :
		fmt.Println("non number")
	}


}