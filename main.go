package main

import "fmt"

func main() {
	name := "nico"
	// 함수 내부에서만 위와 같이 선언이 가능합니다. 또한 첫 번쨰 값을 기준으로 변수의 타입이 결정됩니다.
	// var name string = "nico"
	name = "Park"
	fmt.Println(name)
}
