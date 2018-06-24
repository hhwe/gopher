package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Printf("%p\n", &stu)
		stu.Age = 100
		m[stu.Name] = &stu
	}
	fmt.Println(stus)
	return m
}

func main() {
	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
