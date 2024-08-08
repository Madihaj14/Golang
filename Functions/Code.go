package main

import "fmt"

//Ex 1 Functions
func concat(s1 string, s2 string) string {
	return s1 + s2
}

func mAin() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")

}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
