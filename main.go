package main

import (
	"fmt"

	"github.com/souaremsaidou/webservice/pkg"
)

func main() {

	c := pkg.NewC()
	// fmt.Println(c)
	fmt.Println(c.MethodName())
}
