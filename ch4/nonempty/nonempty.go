package main

import "fmt"

func nonempty(strings []string) []string {
	fmt.Printf("%T\n", strings)
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func remove(slice []int, i int) []int {
	fmt.Println(slice[i:])
	fmt.Println(slice[i+1:])
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	return slice[:len(slice)-1]
}

func main() {
	//data := []string{"one", "", "three"}
	//fmt.Printf("%q\n", nonempty(data))
	s := []int{1, 2, 3, 4, 5}
	s = remove(s, 2)
	fmt.Println(s)

}
