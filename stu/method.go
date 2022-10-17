package main
import "fmt"
func main(){
	g:= greeter{// tạo đối tượng g
		greeting : "hello",
		name: "nguyen",
	}
	q:=greeter{
		greeting: "chao",
		name: "dung",
	}
	g.greet()//gọi method g
	fmt.Println(g.name)//line 18: vẫn là nguyên
	q.qreet()//gọi method q có con trỏ
	fmt.Println(q.name)
}
type greeter struct{
	greeting string
	name string
}
//method là hàm của các đối tượng trong struct
func (g greeter) greet(){//tên hàm : greet
	fmt.Println(g.greeting, g.name)
	g.name = "bach"//k thay đổi
}
func (q *greeter) qreet(){//tên hàm : greet
	fmt.Println(q.greeting, q.name)
	q.name = "bach"//thay đổi
}