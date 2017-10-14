package main

import "fmt"

func main() {
	arr := [5]int{5, 3, 6, 7, 1}
	str := ""
	for i := range arr {
		if i > 0 {
			arr[i] = 0
			str += fmt.Sprint(arr[i])
			if i != 4 {
				str += ","
			}
		} else {
			str += fmt.Sprint(arr[i])
			str += ","
		}
	}
	println(str)
}
