package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretPoop struct {
	person
	ltk bool
}

var mu map[string]int

func main() {
	//Array
	var x [5]int
	fmt.Println(x)
	//Slice
	a := []int{4, 5, 6, 7}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	//Range
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i, v := range a {
		fmt.Println(i, v)
	}
	for _, v := range a {
		fmt.Println(v)
	}
	//Slicing the slice
	fmt.Println(a[1:3])
	fmt.Println(a[:2])
	//Append
	a = append(a, 2, 3, 4, 232324, 434)
	fmt.Println(a)
	y := []int{444, 555, 666}
	a = append(a, y...)
	fmt.Println(a)
	//Delete from the slice
	a = append(a[:3], a[10:]...)
	fmt.Println(a)
	//Making like reserve in cpp
	slice := make([]int, 10, 100)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	//Multi-dimensional slice
	xp := [][]string{{"lolik", "bolik"}, {"noob", "or"}}
	fmt.Println(xp)
	//Map
	m := map[string]int{
		"Lolik":  23,
		"Noobik": 76337,
	}
	fmt.Println(m)
	v, ok := m["Lolik"]
	fmt.Println(v, ok)
	v, ok = m["Noobas"]
	fmt.Println(v, ok)
	if v, ok = m["Noobik"]; ok {
		fmt.Println("OK - ", v)
	}
	if v, ok = m["lolikus"]; ok {
		fmt.Println("should not run")
	}
	//Delete
	delete(m, "Lolik")
	delete(m, "fjdsaghsjdhadsl")
	// and it's not error

	//Struct
	p1 := person{
		first: "Fool",
		last:  "lolik",
	}
	p2 := person{
		first: "poop",
		last:  "Poopster",
	}
	scA := secretPoop{
		person: person{
			first: "Poopli",
			last:  "MEGApoopik",
		},
		ltk: true,
	}
	fmt.Println(p1.first, "\t", p2)
	fmt.Println(scA.person.first, scA.last, "\n", scA)
	//Anonymous struct
	lolik_anonym := struct {
		lolik    bool
		noob_lvl int
	}{
		lolik:    true,
		noob_lvl: 12431349032,
	}
	fmt.Println(lolik_anonym)
	//Function
	foo()
	bar("Arseniy")
	s1 := woo("Lolik")
	fmt.Println(s1)
	oi, oip := lolfunc("lolas", true)
	fmt.Println(oi, oip)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println(s, "is lolik")
}

func woo(s string) string {
	return fmt.Sprint(`Hello,
I'm `, s)
}

func lolfunc(name string, islol bool) (int, bool){
	fmt.Println(islol)
	if v, ok := mu[name]; ok{
		return v, ok
	}else{
		return v, ok
	}
}