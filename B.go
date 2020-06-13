package main

import (
	"fmt"
)
type human struct{
	nickname string
	age int
}

type fool struct{
	human
	cool bool
}

func (s fool) lolik_use(){
	fmt.Println("kfewguighuwehg", s)
}

func (s human) lolik_use(){
	fmt.Println("kfewguighuwehg", s)
}

type humane interface{
	lolik_use()
}

func bar(h humane){
	switch h.(type){
	case human:
		fmt.Println(h.(human).nickname)
	case fool:
		fmt.Println(h.(fool).cool)
	}
}

func main(){
	var a [5]int
	fmt.Println(a)
	v := []int{4, 5, 2}
	y := []int{3, 3, 3, 3}
	v = append(v, 0, 0, 0)
	v = append(v, y...)
	fmt.Println(v)
	for i := 0; i < len(v); i++{
		fmt.Print(v[i], " ")
	}
	fmt.Println()
	for _, val := range v{
		fmt.Print(val, " ")
	}
	fmt.Println()
	v = append(v[0:3], v[4:]...)
	v = append(v[0:3])
	for _, i := range v{
		fmt.Print(i, " ")
	}
	fmt.Println()
	b := make([]int, 10, 100)
	fmt.Println(len(b), cap(b), b)

	m := map[string]int{
		"lolik": 23,
		"bolik": 33,
	}
	if i, ok := m["loliuk"]; ok{
		fmt.Println(i);
	}else{
		fmt.Println("no such name")
	}
	delete(m, "lolik")

	h1 := human{
		nickname: "noobas",
		age: 13,
	}

	h2 := fool{
		human: human{
			nickname: "nool",
			age: 996,
		},
		cool: false,
	}

	h2.lolik_use()
	bar(h1)
	bar(h2)

	fmt.Println(h1)
	fmt.Println(h2)
	defer foop(0, 2, 3, 4,3 ,2)
	opi := []int{2, 4, 3, 4}
	foop(opi...)

	func(x int){
		fmt.Println(x)
	}(42)
}

func foop(x ...int){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}