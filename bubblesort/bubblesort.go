package main

import "fmt"

func Swap(slices []int, index int) {
	slices[index], slices[index+1] = slices[index+1], slices[index]
}
func BubbleSort(intergersList []int) {
	length := len(intergersList)
	for i := 0; i < length ; i++ {
		isSwaped := false;
		for j := 0; j < length - i -1; j++ {
			if intergersList[j] > intergersList[j+1] {
				Swap(intergersList, j)
				isSwaped = true
		}
		
	}
	if !isSwaped {
		break
	}

}
}

func main() {
	slicesIntergers := []int{}
	var input,ok int = 9 , 3
	fmt.Println("Enter up to 10 intergers to sort: ");
	for i :=0; i<10; i++{
		_, err := fmt.Scan(&input)
		if err != nil{
			fmt.Println("Error reading interger")
			break
		}
		slicesIntergers = append(slicesIntergers, input)
	}
	BubbleSort(slicesIntergers)
	fmt.Println(slicesIntergers)

}