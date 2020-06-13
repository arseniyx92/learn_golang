package main

import "fmt"

//func main() {
//	fmt.Println(bart()())
//}
//
//func bart() func() int{
//	return func() int{
//		return 55
//	}
//}

func main(){
	a := []int{3, 4, 5, 2, 5, 2}
	fmt.Println(fu(a...))
	fmt.Println(gu(fu, a...))
}

func fu(a ...int) int{
	summ := 0
	for _, v := range a{
		summ += v
	}
	return summ
}

func gu(f func(a ...int) int, a ...int) int{
	y := []int{}
	for _, v := range a{
		if v%2 == 0{
			y = append(y, v)
		}
	}
	return fu(y...)
}

//117
//func (x int){
//	fmt.Println(x)
//}(32)
//f := func(x int){
//	fmt.Println("lolik", x);
//}
//f(32)