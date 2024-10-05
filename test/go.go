package main

import "fmt"

var funcVar func(int) int

func incFc(x int) int { return x + 1 }
func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	funcVar = incFc       // Đúng: Gán hàm incFc cho biến funcVar
	fmt.Print(funcVar(1)) // Kết quả: 2

	fB := fA()
	fmt.Print(fB())
	fmt.Print(fB())
    
}
