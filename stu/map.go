package main
import "fmt"
func main(){
	//studentNameAgeMap := make(map[string]int) khoi tao neu chua co gia tri
	studentNameAgeMap :=map [string]int{//nameMap := map [typeKey] typeValue
		"Nguyen" : 19,//key : value,
		"Bach" : 20,
		"Dung" : 21,
	}
	studentNameAgeMap["Huy"]=22//them 1 cap
	studentNameAgeMap["Nguyen"]=22//sua 1 cap
	fmt.Println(studentNameAgeMap)//map[bach:20 dung:21 huy:22 nguyen:22]
	fmt.Println(studentNameAgeMap["Nguyen"])//22
	delete(studentNameAgeMap, "Huy")//delete(mapName, "Key")
	fmt.Println(studentNameAgeMap)//map[bach:20 dung:21 nguyen:22]
	ageNguyen := studentNameAgeMap["Nguyen"]//gan gia tri bang key
	fmt.Printf("Tuoi cua nguyen %v\n",ageNguyen)//22
	_, isExist := studentNameAgeMap["Bach"]//_ -> biết có trả về nhưng k dùng
	fmt.Println(isExist)
	fmt.Println(len(studentNameAgeMap))//3
	copyMap := studentNameAgeMap//truyen tham chieu vao copyMap
	fmt.Println(copyMap)//copyMap thay doi -> strudenNameAgeMap thay doi
}