package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("cc")

	unicode_1()
	fmt.Println(unicode_2("가나"))
	fmt.Println(unicode_2("Go 언어"))
	fmt.Println(unicode_2("그럼"))
	fmt.Println(unicode_2("밥묵자"))

	byte_1()
	PlusString()
	ConvertValue()
	Array_1()
	Talk_1()
}

func unicode_1() {

	for i, r := range "가나다" {
		fmt.Println(i, r)
	}

	//range 뒤에 글자를 순차적으로 넣어준다 yield 인가 보다 foreach?
	for _, r := range "가갛힣" {
		fmt.Println(string(r), r)
	}
}

var (
	start = rune(44032)
	end   = rune(55204)
)

func unicode_2(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

func byte_1() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()
}

//문자열 붙이기
func PlusString() {
	s := "abc" //변수선언
	ps := &s   //포인터 할당
	s += "def" //붙이기
	fmt.Println(s)
	fmt.Println(*ps)
}

func ConvertValue() {
	//변환
	var err error
	var num1 int
	num1, err = strconv.Atoi("100") // 문자열 "100"을 숫자 100으로 변환
	fmt.Println(num1, err)          // 100 <nil>

	var num2 int
	num2, err = strconv.Atoi("10t") // 10t는 숫자가 아니므로 에러가 발생함
	fmt.Println(num2, err)          // 0 strconv.ParseInt: parsing "10t": invalid syntax

}

func Array_1() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s 는 맛있다", fruit)
		fmt.Println()
	}
}

func Talk_1() {
	fmt.Println("방갑습니다.")
	fmt.Println("입력해보세요")

	var input string
	fmt.Scanln(&input)
	fmt.Print(input)

	//reader := bufio.NewReader(os.Stdin)
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	//fmt.Println("오예")
}


var task =struct {
	title string
	done bool
	
}