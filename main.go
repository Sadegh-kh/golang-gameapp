package main

import (
	"fmt"
	"gameapp/pkg/hashpassword"
)

func main() {

	testString := "09125jdfhukslhgjklsfdhjkglhsdklghjsdghjksdhjkl3242344"

	fmt.Println(testString[:2])

	passhash:=hashpassword.EncodePasword(testString)
	fmt.Println("password hash: ",passhash)

	pass,_:=hashpassword.DecodePasword(passhash)
	fmt.Println("orginal password: ",pass)

}