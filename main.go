package main

import (
	"encoding/json"
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/hashpassword"
	"gameapp/service/userservice"
	"gameapp/storage/mysql"
	"io"
	"log"
	"net/http"
)

var (
	mySQL       = mysql.New()
	userService = userservice.New(&mySQL)
)

func main() {
	// method 1
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", healthCheckerHandler)
	// mux.HandleFunc("/user/register/", userRegisterHandler)
	// server := http.Server{Addr: ":8080", Handler: mux}
	// log.Println("server in localhost:8080 is listening...")
	// server.ListenAndServe()

	// method 2
	http.HandleFunc("/", healthCheckerHandler)
	http.HandleFunc("/user/register", userRegisterHandler)
	log.Println("server in localhost:8080 is listening...")
	http.ListenAndServe(":8080", nil)

}

func userRegisterHandler(rep http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprintf(rep, `{"error":"invalid method"}`)

		return
	}

	log.Println(`{"log":"request register user post resived"}`)

	data, err := io.ReadAll(req.Body)
	if err != nil {
		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))

		return
	}

	var regReq userservice.RegisterRequest
	err = json.Unmarshal(data, &regReq)
	if err != nil {
		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))

		return
	}

	_, err = userService.Register(regReq)
	if err != nil {
		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))

		return
	}
	rep.Write([]byte(`{"register":"new user created"}`))
}

func healthCheckerHandler(rep http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rep, `{"message":"wellcome to game app"}`)
}

func testMySQLDB() {
	myDb := mysql.New()
	newUser, err := myDb.Register(entity.User{Name: "sadegh", PhoneNumber: "0211412", Password: "3232324"})
	if err != nil {
		fmt.Println("resgister err: ", err)

		return
	}
	fmt.Println("new user:", newUser)

	isUniq, err := myDb.IsPhoneNumberUniq(newUser.PhoneNumber + "213")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("phone number is uniq:", isUniq)
}

func testHashPassword() {

	testString := "09125jdfhukslhgjklsfdhjkglhsdklghjsdghjksdhjkl3242344"

	fmt.Println(testString[:2])

	passhash := hashpassword.EncodePasword(testString)
	fmt.Println("password hash: ", passhash)

	pass, _ := hashpassword.DecodePasword(passhash)
	fmt.Println("orginal password: ", pass)

}
