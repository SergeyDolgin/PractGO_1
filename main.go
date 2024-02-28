// Урок 1. Breakpoints и Простые алгоритмы на массивах

// Задание 1. Слияние отсортированных массивов
// Что нужно сделать:
// Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.

// Задание 2. Сортировка пузырьком
// Что нужно сделать:
// Отсортируйте массив длиной шесть пузырьком.

package main

import "fmt"

func main() {
	array1 := []int{14, 4, 3, 7, 11}
	array2 := []int{21, 13, 55, 7}
	fmt.Println("Исходный первый массив: ", array1)
	fmt.Println("Исходный второй массив: ", array2)
	bubbleSort(array1)
	bubbleSort(array2)
	fmt.Println("Сортированный первый массив: ", array1)
	fmt.Println("Сортированный второй массив: ", array2)
	fmt.Println("Массив, образованный слиянием двух отсортированных массивов: ", union(array1, array2))
	result := union(array1, array2)
	fmt.Println("Общий отсортированный массив: ", result)
	// для задания 2 - сортировка массива длиной шесть пузырьком
	array3 := []int{77, 0, -14, 8, 7, 66}
	fmt.Println("Исходный ассив длиной шесть: ", array1)
	bubbleSort(array3)
	fmt.Println("Отсортированный массив: ", array3)
}

func bubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func union(a []int, b []int) (c []int) {
	var sl []int
	sumArr := sl[:0]
	sumArr = append(sumArr, a[:]...)
	sumArr = append(sumArr, b[:]...)
	sl = sumArr
	return sl
}
