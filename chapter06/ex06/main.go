package main

import "fmt"

func main() {
	// Слайсы могут пересекатсья
	array0 := [5]string{"a", "b", "c", "d", "e"}
	slice01 := array0[0:3]
	slice02 := array0[2:5]
	fmt.Println(slice01, slice02)

	// Изменения в базовом массиве также отразятся и в слайсе
	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice1)

	// Изменения в слайсе также отразятся в базовом массиве
	array2 := [5]string{"a", "b", "c", "d", "e"}
	slice2 := array2[2:5]
	slice2[1] = "X"
	fmt.Println(array2)
	fmt.Println(slice2)

	// Изменения в массиве будут видны во всех его слайсах
	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	array3[2] = "X"
	fmt.Println(array3)
	fmt.Println(slice3, slice4)
}
