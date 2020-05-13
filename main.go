package main

import (
	"fmt"
	"strings"
)

// defer: defer는 함수가 끝난 후(실행되고 난 후)에 추가적으로 무엇인가 동작할 수 있게 하는 기능입니다.
// naked return: return할 변수를 꼭 명시하지 않아도 되는 것을 의미합니다. 이미 어떤 값을 리턴할지를 명시해 두었기 깨문에 length := len(name)으로 코드를 작성하지 않습니다.
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, upperName := lenAndUpper("dd")
	fmt.Println(totalLength, upperName)
}
