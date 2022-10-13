package main
import "fmt"
func main(){
	array := [3]int{1,5,7}//duyet mang va slice
	for index, value := range array{
		fmt.Println(index,value)
	}
	m := map[string ]int{//duyet map
		"Nguyen" :19,
		"Bach": 21,
	}
	for key, value :=range m{
		fmt.Println(key,value)
	}
	
}