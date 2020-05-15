package main

import "fmt"

func main() {
	// map은 자바스크립트의 Object와 유사하지만 다른 자료구조입니다. key,value 형태인데 type이 결정되어 있습니다.
	nico := map[string]string{"name": "nico", "age": "12"}
	for _, value := range nico {
		fmt.Println(value)
	}
}
