package main

import "fmt"

// variable expression: 조건을 체크하기전에 if 조건문 안에 variable을 생성할 수 있는 문법. 외부에서 변수를 사용할 수도 있지만
// 이렇게 코드를 작성하면 다른 개발자가 와서 봤을때 if else문에서만 사용하기위해 변수를 생성했음을 알 수 있다.
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
