package main

import "fmt"

type SalaryCalculator interface { //dinh nghia 1 interface
    CalculateSalary() int//phuong thuc
	//trong phuong thuc chua kieu du lieu de thuc hien phuong thuc
}

type Permanent struct { //2 kieu du lieu implement interface
    empId    int
    basicpay int
    pf       int
}

type Contract struct { 
    empId  int
    basicpay int
}

type Intern struct {
    empId int
    studyFee int
}
func (p Permanent) CalculateSalary() int {//phuong thuc tinh luong
    return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {  
    return c.basicpay
}
func (i Intern) CalculateSalary() int{
    return i.studyFee
}
func totalExpense(s []SalaryCalculator) {
	//slice SalaryCalculator chua' tung kieu du lieu tuong ung voi phuong thuc
    expense := 0
    for _, v:= range s {//s=3
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    iemp1 := Intern{4,2000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1,iemp1}
    totalExpense(employees)
}