package main

import "fmt"

func main() {
	var months = [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}

	summer := months[3:6]
	fmt.Println(len(summer))
	fmt.Println(cap(summer))
	s1 := summer[:6]
	fmt.Println(s1)

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[2:])
	fmt.Println(a)

	s2 := make([]int, 10, 10)
	fmt.Println(s2)
	s2 = append(s2, 1)
	fmt.Println(s2)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
