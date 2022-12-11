package main

import "fmt"

func main() {
	// При обращении к элементу слайса, которому не было присвоено значение,
	// получим нулевое значение для этого типа
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[9], boolSlice[5])

	// Переменная слайса, которой не было присвоено значение,
	// будет содержать значение nil
	var intSlice []int
	var stringSlice []string
	fmt.Printf("intSlice: %#v, stringSlice: %#v\n", intSlice, stringSlice)

	// Значение слайса nil интерпретируется как пустой слайс
	fmt.Println(len(intSlice))
	intSlice = append(intSlice, 27)
	fmt.Printf("intSlice: %#v\n", intSlice)

	// В общем случае не нужно беспокоиться, что используеся - пустой слайс или nil-слайс
	var slice []string
	if len(slice) == 0 {
		slice = append(slice, "first item")
	}
	fmt.Printf("%#v\n", slice)
}
