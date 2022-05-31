package main

import "fmt"

var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	for i := 0; i < arr[len(arr)-1]; i++ {
		fmt.Println(arr[i])
	}
}
