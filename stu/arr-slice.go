package main
import "fmt"
func main(){
	//arr---------------------------------------------------
	i := [4] int {1,2,3,4}
	fmt.Printf("%v,%T\n",i,i)
	var n[3] int
	n[0]=1
	n[2]=3
	fmt.Printf("%v,%v\n",n[0],n[2])//1 3
	fmt.Printf("%v\n",n)//[1 0 3]
	j := [...] int {1,2,3,4}// ... -> tu dem so luong ptu
	//len(arr) : so luong ptu trong mang
	//cap(arr) : suc chua 
	fmt.Printf("%v,%v\n",j,len(j))// 4
	//slice-------------------------------------------------
	m := [8] int {1,4,6,2,12,45,15,2}//mang
	z := [] int {3, 4, 5}//slice
	fmt.Println(z)
	// : toán tử thay đổi con trỏ đầu và con trỏ cuối
	b := m[:]// : lay het phan tu gan vao slice b
	c := m[3:]// : lấy từ ptu thứ 3 gán vào slice c
	d := m[:6]// : lấy dưới ptu thứ 6 gán vào slice d
	e := m[3:6]// : lấy từ 3 - 6 gán vào slice e
	fmt.Printf("m %v,%v,%v\n",m, len(m),cap(m))
	fmt.Printf("b %v,%v,%v\n",b, len(b),cap(b))
	fmt.Printf("c %v,%v,%v\n",c, len(c),cap(c))// 3: = [2 12 45 15 2], 5, 5
	fmt.Printf("d %v,%v,%v\n",d, len(d),cap(d))// :6 = [1 4 6 2 12 45], 6, 8
	fmt.Printf("e %v,%v,%v\n",e, len(e),cap(e))// 3:6 = [2 12 45], 3, 5
    //silce nhu 1 con tro mang-------------------------------------------
	e[2] = 3//e 3:6
	fmt.Printf("%v\n",m)//[1 4 6 2 -> 12 3 15 2]
	//tao slice bang make va append--------------------------------------
	u :=make([]int,0,0)// ham make: tao slice len = 0, cap = 0
	fmt.Printf("u %v,%v,%v\n",u, len(u),cap(u))//u[],0,0
	u = append(u,1)//ham append: them 1 hoac nhieu ptu vao slice, tang kich thuoc len va cap
	u = append(u,4, 6)
	fmt.Printf("u %v,%v,%v\n",u, len(u),cap(u))//[1 4 6] 3, 3
	u= append(u,[]int{2,5}...)//toan tu ... de tham chieu mang vao slice
	fmt.Printf("u %v,%v,%v\n",u, len(u),cap(u))//[1 4 6 2 5] 5, 6
	//stack---------------------------------------------------
	p := []int{1, 2, 3, 4, 5, 6}
	o := p[:5]//pop ptu o dinh: 6
	fmt.Printf("p %v\n",p)//[1 2 3 4 5 6]
	fmt.Printf("o %v\n",o)//[1 2 3 4 5]
	o = append(o, 10)//push ptu vao dinh
	fmt.Printf("o %v\n",o)//[1 2 3 4 5 10]
	fmt.Printf("p %v\n",p)//[1 2 3 4 5 10]
	
}