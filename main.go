package main

import (
	"fmt"
	"strings"
)

// 들어오는 argument의 값이 정해져 있지 않은 경우 이렇게 작성할 수 있습니다.
func repeatMe(words ...string) {
	fmt.Println(words)
}

// GO는 특이하게 함수가 동시에 여러개의 값을 반환할 수 있습니다.
func lenAndUpprt(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// GO는 argument와 return 값의 type을 함수에 명시해야합니다.
func multiply(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
	totalLength, upperName := lenAndUpprt("dd")
	fmt.Println(totalLength, upperName)
	repeatMe("DD", "Hi")
}
