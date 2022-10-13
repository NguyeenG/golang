package main

import 
	"fmt"
type Student struct {
	number int
	name string
	isMale bool
	subject []string
}
type People struct{
	ten string
	age int
}
type Animal struct{
	People//tai su dung struct line 43
	number int 
	food []string
}
func main(){
	studentNguyen := Student{
		number : 3,
		name : "Nguyen",
		isMale: true,
		subject : []string{
			"Math",
			"English",
		},
	}
	fmt.Println(studentNguyen.name)//Nguyen
	fmt.Println(studentNguyen.subject)//[Math English]
	studentBach := struct {name string} {name : "Bach"}//khoi tao struc nhanh
	fmt.Println(studentBach)//Bach
	copyStudent := studentBach//thuc hien copy k bi tham chieu
	copyStudent.name = "Ngu"
	fmt.Println(copyStudent)//Ngu
	fmt.Println(studentBach)//Bach -> k bi thay doi
	//neu muon tham chieu them dau &
	copyStudent1 := &studentNguyen
	copyStudent1.name = "Gioi"
	fmt.Println(copyStudent1.name)//Gioi
	fmt.Println(studentNguyen.name)//Gioi -> da bi thay doi
	dog := Animal{}
	dog.ten = "Cun"//dung duoc thanh phan cua struc People
}