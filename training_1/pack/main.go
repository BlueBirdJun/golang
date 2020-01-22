package main

func main() {
	//변수할당

	//상수
	const c int = 10
	const s string = "Hi"
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)
	println(Visa)
	//타입변환
	var i int = 100
	//var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(f, u)

	//str := "ABC"
	//bytes := []byte(str)
	//str2 := string(bytes)
	//println(bytes, str2)

	//산술연산자
	//관계연산자
	//논리연산자
	//bitwise 연산자
	//swich문

	var name string
	var category = 1
	switch category {
	case 1:
		name = "peter book"
	case 2:
		name = "ebook"
	case 3:
		name = "blog"
	default:
		name = "other"
	}
	println(name)

	sum := 0
	for d := 1; d <= 10; d++ {
		sum += d
	}
	println(sum)

}
