package main

import (
	"fmt"

	greetingsv1 "github.com/handrixn/gohellopackage/greetings"
	"github.com/handrixn/gohellopackage/v2/greetings"
)

func main() {
	str1 := greetings.SayHello("test")
	str2 := greetingsv1.SayHello()
	fmt.Println(str1)
	fmt.Println(str2)
}
