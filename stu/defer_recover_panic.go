package main
import "fmt"
func main(){
	//defer--------------------------------------
	fmt.Println("Line 1")
	defer fmt.Println("Line 2")//bỏ qua lệnh này đến cuối func mới chạy
	fmt.Println("Line 3")
	//line 1
	//line 3
	//line 2 *
	//defer dùng stack -> vào sau ra trước
	a := 12
	defer fmt.Println(a)
	a = 24
	//12 : defer dua a vao stack luc a= 12
	//panic------------------------------------------
	//n :=10
	//m:=0 -> vi du nguoi dung nhap tu ban phim
	//fmt.Println(n/m)k the chia cho 0
	//luc nay xuat hien loi panic
	//dung luc du doan duoc se co loi
	defer fmt.Println("Hello")//van chay duoc
	defer func ()  {//thong bao cho nguoi dung lam gi se co loi
		if err := recover();err !=nil{
			fmt.Println("Error: ",err)
		}	
	}()
	panic("Chia 1 so la 0")//thong bao cho nguoi dung luc chay den dong nay ma bi loi
	fmt.Println("hello")//k chay duoc -> tao ham panicker
}
func panicker(){
	defer func(){
	if error := recover();error !=nil{
		fmt.Println("Error: ", error)
	}
}()
	panic("Chia cho 1 so la 0")
}//truyen ham nay vao main de bao loi xong tiep tuc chay nhung dong lenh o phia duoi
