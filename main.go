package main

import (
	"fmt"
	"gameapp/entity"
	"gameapp/storage/mysql"
)

func main() {

	// testString := "09125jdfhukslhgjklsfdhjkglhsdklghjsdghjksdhjkl3242344"

	// fmt.Println(testString[:2])

	// passhash:=hashpassword.EncodePasword(testString)
	// fmt.Println("password hash: ",passhash)

	// pass,_:=hashpassword.DecodePasword(passhash)
	// fmt.Println("orginal password: ",pass)

	myDb:=mysql.New()
	newUser,err:=myDb.Register(entity.User{Name: "sadegh",PhoneNumber: "0421412",Password: "3232324"})
	if err!=nil{
		fmt.Println("resgister err: ",err)
		
		return
	}
	fmt.Println("new user: ",newUser)

}