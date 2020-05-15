package main

import "fmt"

func main() {
	// array는 크기가 정해져 있는 배열, 배열 크기를 제한하지 않고 싶자면 slice를 사용
	// names := [5]string{"nico", "lynn", "dal"}
	// names[3] = "lalala"
	// fmt.Println(names[4])

	names := []string{"nico", "lynn", "dal"}
	// names[3] = "lalala"
	// append는 names을 수정하지 않고 새로운 변수가 추가된 slice를 반환합니다.
	names = append(names, "lalala")
	fmt.Println(names)
}
