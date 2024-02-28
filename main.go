package main

import (
	"gameapp/config"
	"gameapp/delivery/httpserver"
	"gameapp/service/authservice"
	"gameapp/service/userservice"
	"gameapp/storage/mysql"
	"gameapp/validator/uservalidator"
	"time"
)

const (
	SecretKey            = "secret"
	AccessTokenDuration  = time.Hour * 24
	RefreshTokenDuration = time.Hour * 24 * 7
)

type Services struct {
	User          userservice.Service
	Auth          authservice.Service
	UserValidator uservalidator.Validator
}

func main() {
	// method 1
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", healthCheckerHandler)
	// mux.HandleFunc("/user/register/", userRegisterHandler)
	// server := http.Server{Addr: ":8080", Handler: mux}
	// log.Println("server in localhost:8080 is listening...")
	// server.ListenAndServe()

	// method 2
	//http.HandleFunc("/", healthCheckerHandler)
	//http.HandleFunc("/user/register", userRegisterHandler)
	//http.HandleFunc("/user/login", userLoginHandler)
	//http.HandleFunc("/user/profile", profileHandler)
	//log.Println("server in localhost:8080 is listening...")
	//http.ListenAndServe(":8080", nil)

	// method 3 (echo framework)
	cfg := config.Config{
		HttpConf: config.HttpConfig{
			Port: 8999,
		},
		Auth: authservice.Config{
			SecretKey:            SecretKey,
			AccessTokenDuration:  AccessTokenDuration,
			RefreshTokenDuration: RefreshTokenDuration,
			AccessSubject:        "ac",
			RefreshSubject:       "rf",
		},
		MySQL: mysql.Config{
			User:         "gameapp",
			Password:     "gameappt0lk2o20@",
			Address:      "localhost",
			Port:         3310,
			DataBaseName: "gameapp_db",
		}}
	//migrator
	//mig := migrator.New(cfg.MySQL)
	//mig.Up()

	services := setupService(cfg)

	server := httpserver.New(cfg, services.Auth, services.User, services.UserValidator)
	server.Serve()

}

func setupService(cfg config.Config) Services {
	authS := authservice.New(cfg.Auth)
	mySQL := mysql.New(cfg.MySQL)
	uValidator := uservalidator.New(&mySQL)
	userS := userservice.New(&mySQL, authS)

	return Services{
		User:          userS,
		Auth:          authS,
		UserValidator: uValidator,
	}
}

//func profileHandler(rep http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodGet {
//		// fprint called rep's write method
//		fmt.Fprintln(rep, `{"error":"invalid method"}`)
//
//		return
//	}
//
//	log.Println(`{"log":"request profile Get resived"}`)
//
//	authToken := req.Header.Get("Authorization")
//	authToken = strings.Replace(authToken, "Bearer ", "", 1)
//
//	uid, err := authService.ParseToken(authToken)
//	if err != nil {
//		fmt.Fprintf(rep, `{"error":"%v"}`, err)
//
//		return
//	}
//	proReq := userservice.ProfileRequest{UserID: uid}
//
//	proRep, err := userService.Profile(proReq)
//	if err != nil {
//		fmt.Fprintf(rep, `{"error":"%v"}`, err)
//
//		return
//	}
//
//	fmt.Fprintf(rep, `{"user_name":%s}`, proRep.Name)
//
//}

//func userLoginHandler(rep http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodPost {
//		fmt.Fprintf(rep, `{"error":"invalid method"}`)
//
//		return
//	}
//
//	log.Println(`{"log":"request login user post resived"}`)
//
//	data, err := io.ReadAll(req.Body)
//	if err != nil {
//		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//
//		return
//	}
//
//	var logReq userservice.LoginRequest
//	err = json.Unmarshal(data, &logReq)
//	if err != nil {
//		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//
//		return
//	}
//
//	token, err := userService.Login(logReq)
//	if err != nil {
//		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//
//		return
//	}
//
//	repLog, err := json.Marshal(token)
//	if err != nil {
//		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//
//		return
//	}
//	rep.Write(repLog)
//}

//func userRegisterHandler(rep http.ResponseWriter, req *http.Request) {
//	if req.Method != http.MethodPost {
//		fmt.Fprintf(rep, `{"error":"invalid method"}`)
//
//		return
//	}
//
//	log.Println(`{"log":"request register user post resived"}`)
//
//	data, err := io.ReadAll(req.Body)
//	if err != nil {
//		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//
//		return
//	}
//
//	var regReq userservice.RegisterRequest
//	err = json.Unmarshal(data, &regReq)
//	if err != nil {
//		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//
//		return
//	}
//
//	_, err = userService.Register(regReq)
//	if err != nil {
//		rep.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err)))
//
//		return
//	}
//	rep.Write([]byte(`{"register":"new user created"}`))
//}

//func healthCheckerHandler(rep http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(rep, `{"message":"wellcome to game app"}`)
//}
//
//func testMySQLDB() {
//	myDb := mysql.New()
//	newUser, err := myDb.Register(entity.User{Name: "sadegh", PhoneNumber: "0211412", Password: "3232324"})
//	if err != nil {
//		fmt.Println("resgister err: ", err)
//
//		return
//	}
//	fmt.Println("new user:", newUser)
//
//	isUniq, err := myDb.IsPhoneNumberUniq(newUser.PhoneNumber + "213")
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("phone number is uniq:", isUniq)
//}
//
//func testHashPassword() {
//
//	testString := "09125jdfhukslhgjklsfdhjkglhsdklghjsdghjksdhjkl3242344"
//
//	fmt.Println(testString[:2])
//
//	passhash := hashpassword.EncodePasword(testString)
//	fmt.Println("password hash: ", passhash)
//
//	pass, _ := hashpassword.DecodePasword(passhash)
//	fmt.Println("orginal password: ", pass)
//
//}
