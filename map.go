package main
import "fmt"

func main(){

	//create empty map
	var angka = make(map[string]int) // empty map

	//we can input value to empty map
	angka["satu"] = 1
	angka["dua"] = 2
	angka["tiga"] = 3

	fmt.Println(angka["satu"])


	//example 1 map with value
	var indonesia = map[string]string{
		"jawa_tengah" : "semarang",
		"jawa_timur" : "surabaya",
		"jawa_barat" : "bandung",
		"kalsel" : "samarinda",
		"amerika" : "new_york",
	}

	//change value of map
	indonesia["kalsel"] = "banjarmasin"

	//remove value of map
	delete(indonesia, "amerika")

	//check spesific element in map
	provinsi1, ibukota1 := indonesia["kalsel"]
	fmt.Println(provinsi1, ibukota1)

}