package main

import "fmt"

// 일반적인 switch문과 크게 다르지 않습니다. if문과 마찬가지로 variavle expression을 사용할 수 있습니다.
func canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge >= 20 {
	case true:
		return true
	case false:
		return false
	}
	return false
}

func main() {
	fmt.Println(canIDrink(17))
}
