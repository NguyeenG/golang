package main
import "fmt"
var tc int = 10
var Tc int = 12 // viet hoa chu cai dau cho phep package khac
func main(){
	fmt.Printf("%v,%T\n",tc,tc)// tc = 10
	tc = 12
	fmt.Printf("%v,%T\n",tc,tc)// bien tc bi thay doi
}