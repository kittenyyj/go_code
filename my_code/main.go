package main

import (
	"fmt"
	"strings"
)

type Stu struct {
	Name string
	age  []int
}

func main() {
	s := ",hahh,"
	res := strings.Split(s, ",")
	fmt.Println(len(res))
	var sl []int
	sl = append(sl, 1)
	fmt.Println(sl, len(sl), cap(sl))
	f(11, 22, 33)
	fmt.Println()
	str := "1我们1"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c",str[i])
	// }
	for _, v := range str {
		fmt.Printf("%c", v)
	}
	fmt.Println()
	ff := 5 / 4.0
	fmt.Println(ff)
	fmt.Printf("%T", ff)

}
func f(a ...interface{}) {
	fmt.Printf("%T", a)
	for k, v := range a {
		fmt.Println(k, v)
	}
	fmt.Print("1111")
	fmt.Print("44444")
	fmt.Print("22222")
}
